package mapimpl

import "sync"

type RWMutexMap struct {
	Map map[any]any
	mu  sync.RWMutex
}

func NewRWMutex() *RWMutexMap {
	return &RWMutexMap{
		Map: make(map[any]any),
		mu:  sync.RWMutex{},
	}
}

func (m *RWMutexMap) Store(key any, val any) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Map[key] = val
}

func (m *RWMutexMap) Load(key any) (any, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.Map[key]
	return v, ok
}
