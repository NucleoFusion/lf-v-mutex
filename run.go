package main

import (
	"lf-v-mutex/benchmark"
	"lf-v-mutex/mapimpl"

	"github.com/zhangyunhao116/skipmap"
)

func Run() Results {
	return Results{
		LockFree:        RunLockfree(50),
		RWMutex:         RunRWMutex(50),
		PartialLockfree: RunPartialLockfree(50),
	}
}

func RunLockfree(n int) []Result {
	items := make([]Result, 0, n)
	for i := 0; i < n; i++ {
		m := skipmap.New[string, int]()
		r := Result{
			Balanced:   mapimpl.Simulate(m, 1000, 10000, benchmark.Balanced, benchmark.LockFree),
			ReadHeavy:  mapimpl.Simulate(m, 1000, 10000, benchmark.ReadHeavy, benchmark.LockFree),
			WriteHeavy: mapimpl.Simulate(m, 1000, 10000, benchmark.WriteHeavy, benchmark.LockFree),
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
			Balanced:   mapimpl.Simulate(m, 1000, 10000, benchmark.Balanced, benchmark.RWMut),
			ReadHeavy:  mapimpl.Simulate(m, 1000, 10000, benchmark.ReadHeavy, benchmark.RWMut),
			WriteHeavy: mapimpl.Simulate(m, 1000, 10000, benchmark.WriteHeavy, benchmark.RWMut),
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
			Balanced:   mapimpl.Simulate(m, 1000, 10000, benchmark.Balanced, benchmark.PartialLockfree),
			ReadHeavy:  mapimpl.Simulate(m, 1000, 10000, benchmark.ReadHeavy, benchmark.PartialLockfree),
			WriteHeavy: mapimpl.Simulate(m, 1000, 10000, benchmark.WriteHeavy, benchmark.PartialLockfree),
		}
		items = append(items, r)
	}

	return items
}
