package mapimpl

import "sync"

type RWMutexMap struct {
	Map map[string]int
	mu  sync.RWMutex
}

func NewRWMutex() *RWMutexMap {
	return &RWMutexMap{
		Map: make(map[string]int),
		mu:  sync.RWMutex{},
	}
}

func (m *RWMutexMap) Store(key string, val int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Map[key] = val
}

func (m *RWMutexMap) Load(key string) (int, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.Map[key]
	return v, ok
}
