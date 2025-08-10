package mapimpl

import "sync"

type PartialLockfree struct {
	m *sync.Map
}

func NewPartialLockfree() *PartialLockfree {
	return &PartialLockfree{m: &sync.Map{}}
}

func (s *PartialLockfree) Load(key string) (int, bool) {
	val, ok := s.m.Load(key)
	if !ok {
		return 0, false
	}
	intVal, ok := val.(int)
	return intVal, ok
}

func (s *PartialLockfree) Store(key string, value int) {
	s.m.Store(key, value)
}
