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
	buckets atomic.Pointer[[262_144]*AtomicEntry] // 2^18
}

func NewAtomic() *AtomicMap {
	arr := new([262_144]*AtomicEntry)
	m := AtomicMap{}
	m.buckets.Store(arr)

	return &m
}

func HashKey(k string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(k))
	return h.Sum32() % 262_144
}

func (m *AtomicMap) Load(key string) (int, bool) {
	idx := HashKey(key)
	entryPtr := (*m.buckets.Load())[idx]

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
	for {
		oldptr := m.buckets.Load()
		newptr := *oldptr
		newptr[HashKey(key)] = &AtomicEntry{key: key, value: val}
		if m.buckets.CompareAndSwap(oldptr, &newptr) {
			return
		}
		// Else continue
	}
}
