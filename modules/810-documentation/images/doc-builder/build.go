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
	"github.com/flant/doc_builder/pkg/hugo"
	"k8s.io/klog/v2"
	"net/http"
)

func newBuildHandler(contentDir string) *buildHandler {
	return &buildHandler{
		rootDir: contentDir,
	}
}

type buildHandler struct {
	rootDir string
}

func (b *buildHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	err := b.build()
	if err != nil {
		klog.Error(err)
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (b *buildHandler) build() error {
	flags := hugo.Flags{
		Quiet:  true,
		Source: b.rootDir,
	}

	return hugo.Build(flags)
}