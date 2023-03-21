// Copyright 2022 Flant JSC
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
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"

	"discoverer/app"
	"discoverer/internal"
	"discoverer/internal/openstack"
)

func main() {
	kpApp := kingpin.New("openstack cloud data discoverer", "A tool for discovery data from cloud provider")
	kpApp.HelpFlag.Short('h')

	app.InitFlags(kpApp)

	kpApp.Action(func(context *kingpin.ParseContext) error {
		logger := app.InitLogger()
		client := app.InitClient(logger)
		discoverer := openstack.NewDiscoverer(logger)

		r := internal.NewReconciler(discoverer, app.ListenAddress, app.DiscoveryPeriod, logger, client)
		r.Start()

		return nil
	})

	_, err := kpApp.Parse(os.Args[1:])
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
