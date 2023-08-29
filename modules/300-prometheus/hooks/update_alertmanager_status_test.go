/*
Copyright 2023 Flant JSC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package hooks

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/deckhouse/deckhouse/testing/hooks"
)

const (
	alertmanagersValues = `{
		"byAddress":[{
			"basicAuth": {},
			"name": "my-fqdn-alertmanager",
			"path": "/myprefix",
			"scheme": "https",
			"target": "alertmanager.mycompany.com",
			"tlsConfig": {}
		}],
		"byService": [{
			"name": "alerts-receiver",
			"namespace": "d8-monitoring",
			"pathPrefix": "/",
			"port": "http",
			"resourceName": "alerts-receiver"
		}],
		"internal": [{
			"name": "webhook",
			"receivers": [{
				"name": "webhook",
				"webhookConfigs": [{
					"url": "http://webhookserver:8080/"
				}]
			}],
			"route": {
				"groupBy": ["job"],
				"groupInterval": "5m",
				"groupWait": "30s",
				"receiver": "webhook",
				"repeatInterval: 12h"
			}
		}]
	}`
	nowTime = "2023-03-03T16:49:52Z"
)

var _ = Describe("Modules :: prometheus :: hooks :: update alertmanagers' statuses", func() {
	f := HookExecutionConfigInit(`{"prometheus": {"internal": {}}}`,
		`{"prometheus":{}}`,
	)
	f.RegisterCRD("deckhouse.io", "v1alpha1", "CustomAlertmanager", false)

	err := os.Setenv("TEST_CONDITIONS_CALC_NOW_TIME", nowTime)
	if err != nil {
		panic(err)
	}

	Context("Alertmanagers' statuses are updated", func() {
		BeforeEach(func() {
			f.ValuesSetFromYaml("prometheus.internal.alertmanagers", []byte(alertmanagersValues))
			f.KubeStateSet(testAlertmanagers)
			f.BindingContexts.Set(f.GenerateAfterHelmContext())
			f.RunHook()
		})
		It("should have generated resources", func() {
			Expect(f).To(ExecuteSuccessfully())
			Expect(f.ValuesGet("prometheus.internal.alertmanagers.byAddress").Array()).To(HaveLen(1))
			Expect(f.ValuesGet("prometheus.internal.alertmanagers.byService").Array()).To(HaveLen(1))
			Expect(f.ValuesGet("prometheus.internal.alertmanagers.internal").Array()).To(HaveLen(1))
			const expectedStatus = `{
				"deckhouse": {
        				"processed": {
						"generation": 0,
						"lastTimestamp": "2023-03-03T16:49:52Z"
					}
				}
			}`
			Expect(f.KubernetesGlobalResource("CustomAlertmanager", "alerts-receiver").Field("status").String()).To(MatchJSON(expectedStatus))
			Expect(f.KubernetesGlobalResource("CustomAlertmanager", "my-fqdn-alertmanager").Field("status").String()).To(MatchJSON(expectedStatus))
			Expect(f.KubernetesGlobalResource("CustomAlertmanager", "webhook").Field("status").String()).To(MatchJSON(expectedStatus))
		})
	})
})

var testAlertmanagers = `
---
apiVersion: deckhouse.io/v1alpha1
kind: CustomAlertmanager
metadata:
  generation: 1
  name: alerts-receiver
spec:
  external:
    service:
      name: alerts-receiver
      namespace: d8-monitoring
      path: /
  type: External
---
apiVersion: deckhouse.io/v1alpha1
kind: CustomAlertmanager
metadata:
  name: my-fqdn-alertmanager
spec:
  external:
    address: https://alertmanager.mycompany.com/myprefix
  type: External
---
apiVersion: deckhouse.io/v1alpha1
kind: CustomAlertmanager
metadata:
  name: webhook
spec:
  internal:
    receivers:
    - name: webhook
      webhookConfigs:
      - url: http://webhookserver:8080/
    route:
      groupBy:
      - job
      groupInterval: 5m
      groupWait: 30s
      receiver: webhook
      repeatInterval: 12h
  type: Internal
`
