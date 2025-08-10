package mapimpl

import (
	"hash/fnv"
	"sync/atomic"
)

type AtomicEntry struct {
	key   string
	value int
}

type AtomicMap struct {
	buckets [262_144]atomic.Pointer[AtomicEntry] // 2^18
}

func NewAtomic() *AtomicMap {
	return &AtomicMap{}
}

func HashKey(k string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(k))
	return h.Sum32() % 262_144
}

func (m *AtomicMap) Load(key string) (int, bool) {
	idx := HashKey(key)
	entryPtr := m.buckets[idx].Load()

	// Not found
	if entryPtr == nil {
		return 0, false
	}
	// Found
	if entryPtr.key == key {
		return entryPtr.value, true
	}
	return 0, false // Poor hash collision handling
}

func (m *AtomicMap) Store(key string, val int) {
	idx := HashKey(key)
	entry := &AtomicEntry{key: key, value: val}
	m.buckets[idx].Store(entry)
}
