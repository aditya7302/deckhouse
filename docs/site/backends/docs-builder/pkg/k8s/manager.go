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

package k8s

import (
	"context"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/utils/pointer"
	"os"
	"strings"
	"time"

	coordination "k8s.io/api/coordination/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ns    = "d8-system"
	name  = "docs-sync"
	label = "deckhouse.io/documentation-builder-sync"

	leaseDuration = 35
)

func NewLeasesManager() (*LeasesManager, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("create rest config: %w", err)
	}

	kclient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("new client set: %w", err)
	}

	return &LeasesManager{kclient: kclient}, nil
}

type LeasesManager struct {
	name string

	kclient *kubernetes.Clientset
}

func (m *LeasesManager) Create(ctx context.Context) error {
	l := newLease()
	l, err := m.kclient.CoordinationV1().Leases(ns).Create(ctx, l, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}

	m.name = l.Name
	return nil
}

func (m *LeasesManager) Run(ctx context.Context) error {
	ticker := time.NewTicker(leaseDuration - 5)

	for {
		select {
		case <-ticker.C:
			err := m.renew(ctx)
			if err != nil {
				return fmt.Errorf("renew lease: %w", err)
			}

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (m *LeasesManager) renew(ctx context.Context) error {
	lease, err := m.kclient.CoordinationV1().Leases(ns).Get(ctx, m.name, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("get lease %q: %w", m.name, err)
	}

	_, err = m.kclient.CoordinationV1().Leases(ns).Update(ctx, lease, metav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("update lease %q: %w", m.name, err)
	}

	return nil
}

func (m *LeasesManager) Remove() error {
	if m.name == "" {
		return nil
	}

	err := m.kclient.CoreV1().ConfigMaps(ns).Delete(context.Background(), m.name, metav1.DeleteOptions{})
	m.name = ""
	return err
}

func newLease() *coordination.Lease {
	address := fmt.Sprintf(
		"%s.%s.pod.cluster.local",
		strings.ReplaceAll(os.Getenv("POD_IP"), ".", "-"),
		os.Getenv("POD_NAMESPACE"),
	)

	return &coordination.Lease{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Lease",
			APIVersion: "coordination.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: name,
			Labels:       map[string]string{label: ""},
		},
		Spec: coordination.LeaseSpec{
			HolderIdentity:       pointer.String(address),
			LeaseDurationSeconds: pointer.Int32(leaseDuration),
		},
	}
}
