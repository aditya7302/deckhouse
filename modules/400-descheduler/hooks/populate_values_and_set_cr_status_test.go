/*
Copyright 2022 Flant JSC

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
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/yaml"

	. "github.com/deckhouse/deckhouse/testing/hooks"
)

const (
	deschedulerCR = `---
apiVersion: deckhouse.io/v1alpha1
kind: Descheduler
metadata:
  name: test
spec:
  deploymentTemplate:
    nodeSelector:
      test: test
  deschedulerPolicy:
    globalParameters:
      evictFailedBarePods: true
    strategies:
      lowNodeUtilization:
        enabled: true
`

	deschedulerDeployment = `---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: descheduler-test
  namespace: d8-descheduler
  labels:
    app: descheduler
    name: test
status:
  replicas: 1
  readyReplicas: 1
`
)

var _ = Describe("Modules :: descheduler :: hooks :: populate_values_and_set_cr_status ::", func() {
	f := HookExecutionConfigInit(`{"descheduler":{"internal":{}}}`, ``)
	f.RegisterCRD("deckhouse.io", "v1alpha1", "Descheduler", false)

	Context("Cluster with descheduler object", func() {
		BeforeEach(func() {
			f.KubeStateSet(deschedulerCR)
			f.BindingContexts.Set(f.GenerateBeforeHelmContext())
			f.RunHook()
		})

		It("Should set values appropriately", func() {
			Expect(f).To(ExecuteSuccessfully())
			var obj map[string]interface{}
			Expect(yaml.Unmarshal([]byte(deschedulerCR), &obj)).Should(Succeed())

			Expect(f.ValuesGet("descheduler.internal.deschedulers.0").Value()).To(BeEquivalentTo(obj))
			Expect(f.ValuesGet("descheduler.internal.deschedulers.1").Exists()).To(BeFalse())
		})
	})

	Context("Set status", func() {
		BeforeEach(func() {
			f.KubeStateSet(deschedulerCR + deschedulerDeployment)
			f.BindingContexts.Set(f.GenerateBeforeHelmContext())
			f.RunHook()
		})

		It("Should set Descheduler CR status appropriately", func() {
			Expect(f.KubernetesGlobalResource("Descheduler", "test").Field("status.ready").Bool()).To(BeTrue())
		})
	})
})
