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

package storage

import "sync"

const MemoryValuesStorageDriver = `memory`

type MemoryValuesStorage struct {
	sync.RWMutex
	values map[string]interface{}
}

func NewMemoryValuesStorage() *MemoryValuesStorage {
	return &MemoryValuesStorage{
		values: make(map[string]interface{}),
	}
}

func (m *MemoryValuesStorage) Set(key string, value interface{}) error {
	defer unlock(m.wlock())
	m.values[key] = value
	return nil
}

func (m *MemoryValuesStorage) Remove(key string) error {
	defer unlock(m.wlock())
	delete(m.values, key)
	return nil
}

func (m *MemoryValuesStorage) Get(key string) (interface{}, bool, error) {
	defer unlock(m.rlock())
	v, ok := m.values[key]
	return v, ok, nil
}

func (m *MemoryValuesStorage) wlock() func() {
	m.Lock()
	return func() { m.Unlock() }
}

// rlock locks mem for reading
func (m *MemoryValuesStorage) rlock() func() {
	m.RLock()
	return func() { m.RUnlock() }
}

func unlock(fn func()) { fn() }
