package main

import (
	"fmt"
	"sync"

	"lf-v-mutex/benchmark"
	"lf-v-mutex/mapimpl"
)

func main() {
	fmt.Printf("\n-----------------Lock-Free(sync.Map)----------------------\n\n")

	// Lock-Free (sync.map)
	m := sync.Map{}
	mapimpl.Simulate(&m, 1000, 10000, benchmark.Balanced, benchmark.LockFree).Print()
	mapimpl.Simulate(&m, 1000, 10000, benchmark.ReadHeavy, benchmark.LockFree).Print()
	mapimpl.Simulate(&m, 1000, 10000, benchmark.WriteHeavy, benchmark.LockFree).Print()

	fmt.Printf("\n----------------------RWMutex-----------------------------\n\n")
	// RWMutex (sync.RWMutex)
	rwm := mapimpl.NewRWMutex()
	mapimpl.Simulate(rwm, 1000, 10000, benchmark.Balanced, benchmark.RWMut).Print()
	mapimpl.Simulate(rwm, 1000, 10000, benchmark.ReadHeavy, benchmark.RWMut).Print()
	mapimpl.Simulate(rwm, 1000, 10000, benchmark.WriteHeavy, benchmark.RWMut).Print()

	fmt.Printf("\n---------------------AtomicMap----------------------------\n\n")
	// Atomic Map(sync/atomic)
	atm := mapimpl.NewAtomic()
	mapimpl.Simulate(atm, 1000, 10000, benchmark.Balanced, benchmark.Atomic).Print()
	mapimpl.Simulate(atm, 1000, 10000, benchmark.ReadHeavy, benchmark.Atomic).Print()
	mapimpl.Simulate(atm, 1000, 10000, benchmark.WriteHeavy, benchmark.Atomic).Print()
}
