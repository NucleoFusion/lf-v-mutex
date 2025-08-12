package main

import (
	"lf-v-mutex/benchmark"
	"lf-v-mutex/mapimpl"

	"github.com/zhangyunhao116/skipmap"
)

func Run() Results {
	return Results{
		LockFree:        RunLockfree(10),
		AtomicMap:       RunAtomic(10),
		RWMutex:         RunRWMutex(10),
		PartialLockfree: RunPartialLockfree(10),
	}
}

func RunLockfree(n int) []Result {
	items := make([]Result, 0, n)
	for i := 0; i < n; i++ {
		m := skipmap.New[string, int]()
		r := Result{
			Balanced:   mapimpl.Simulate(m, 100, 10000, benchmark.Balanced, benchmark.LockFree),
			ReadHeavy:  mapimpl.Simulate(m, 100, 10000, benchmark.ReadHeavy, benchmark.LockFree),
			WriteHeavy: mapimpl.Simulate(m, 100, 10000, benchmark.WriteHeavy, benchmark.LockFree),
		}
		items = append(items, r)
	}

	return items
}

func RunRWMutex(n int) []Result {
	items := make([]Result, 0, n)
	for i := 0; i < n; i++ {
		m := mapimpl.NewRWMutex()
		r := Result{
			Balanced:   mapimpl.Simulate(m, 100, 10000, benchmark.Balanced, benchmark.RWMut),
			ReadHeavy:  mapimpl.Simulate(m, 100, 10000, benchmark.ReadHeavy, benchmark.RWMut),
			WriteHeavy: mapimpl.Simulate(m, 100, 10000, benchmark.WriteHeavy, benchmark.RWMut),
		}
		items = append(items, r)
	}

	return items
}

func RunAtomic(n int) []Result {
	items := make([]Result, 0, n)
	for i := 0; i < n; i++ {
		m := mapimpl.NewAtomic()
		r := Result{
			Balanced:   mapimpl.Simulate(m, 100, 10000, benchmark.Balanced, benchmark.Atomic),
			ReadHeavy:  mapimpl.Simulate(m, 100, 10000, benchmark.ReadHeavy, benchmark.Atomic),
			WriteHeavy: mapimpl.Simulate(m, 100, 10000, benchmark.WriteHeavy, benchmark.Atomic),
		}
		items = append(items, r)
	}

	return items
}

func RunPartialLockfree(n int) []Result {
	items := make([]Result, 0, n)
	for i := 0; i < n; i++ {
		m := mapimpl.NewPartialLockfree()
		r := Result{
			Balanced:   mapimpl.Simulate(m, 100, 10000, benchmark.Balanced, benchmark.PartialLockfree),
			ReadHeavy:  mapimpl.Simulate(m, 100, 10000, benchmark.ReadHeavy, benchmark.PartialLockfree),
			WriteHeavy: mapimpl.Simulate(m, 100, 10000, benchmark.WriteHeavy, benchmark.PartialLockfree),
		}
		items = append(items, r)
	}

	return items
}
