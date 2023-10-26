// Copyright 2023 Flant JSC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"errors"
	"flag"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/flant/docs_builder/pkg/k8s"

	"k8s.io/klog/v2"
)

// flags
var (
	listenAddress string
	src           string
	dst           string
)

func init() {
	flag.StringVar(&listenAddress, "address", ":8081", "Address to listen on")
	flag.StringVar(&src, "src", "/tmp/src", "Directory to load source files")
	flag.StringVar(&dst, "dst", "/tmp/dst", "Directory for compiled files")
}

func main() {
	flag.Parse()

	ctx, stopNotify := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stopNotify()

	cmManager, err := k8s.NewConfigmapManager()
	if err != nil {
		klog.Fatalf("new cm manager: %s", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) { _, _ = w.Write([]byte("ok")) })
	r.Handle("/loadDocArchive/{moduleName}/{version}", newLoadHandler(src)).Methods(http.MethodPost)
	r.Handle("/build", newBuildHandler(src, dst, cmManager)).Methods(http.MethodPost)

	srv := &http.Server{
		Addr:    listenAddress,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			klog.Fatalf("listen: %s", err)
		}
	}()
	klog.Info("Server Started")

	err = cmManager.Create(ctx)
	if err != nil {
		klog.Fatalf("create sync config map: %s", err)
	}

	<-ctx.Done()
	klog.Info("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		klog.Fatalf("Server Shutdown Failed:%+v", err)
	}
}
