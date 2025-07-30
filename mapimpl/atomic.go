package mapimpl

import (
	"hash/fnv"
	"sync/atomic"
)

type AtomicEntry struct {
	key   any
	value any
}

type AtomicMap struct {
	buckets [262_144]atomic.Pointer[AtomicEntry] // 2^18
}

func NewAtomic() *AtomicMap {
	return &AtomicMap{}
}

func HashKey(k any) uint32 {
	h := fnv.New32a()
	if str, ok := k.(string); ok { // Will always happen, to accomodate sync.Map
		h.Write([]byte(str))
	}
	return h.Sum32() % 262_144
}

func (m *AtomicMap) Load(key any) (any, bool) {
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

func (m *AtomicMap) Store(key any, val any) {
	idx := HashKey(key)
	entry := &AtomicEntry{key: key, value: val}
	m.buckets[idx].Store(entry)
}
