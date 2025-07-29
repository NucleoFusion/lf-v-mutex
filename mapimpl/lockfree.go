package mapimpl

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"lf-v-mutex/benchmark"
)

func SimulateLF(numRoutines int, opsPerRoutine int) *benchmark.MetricResult {
	m := sync.Map{}
	var wg sync.WaitGroup

	start := time.Now()

	// Declaring mem stats and Initializing Current stats
	var memBefore, memAfter runtime.MemStats
	runtime.ReadMemStats(&memBefore)

	numOps := numRoutines * opsPerRoutine
	opTimes := make([]int64, 0, numOps)
	cpuStart := cpuTime()

	wg.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go func(id int) {
			defer wg.Done()

			for j := 0; j < opsPerRoutine; j++ {
				t0 := time.Now().UnixNano()

				key := fmt.Sprintf("r%d_op%d", id, j)
				m.Store(key, j) // simulate work

				delta := time.Now().UnixNano() - t0
				opTimes = append(opTimes, delta)
			}
		}(i)
	}
	wg.Wait()

	totalTime := time.Since(start)
	throughput := float64(numOps) / totalTime.Seconds()
	avgLatency := float64(sum(opTimes)) / float64(numOps)

	// Memory and CPU stats after
	runtime.ReadMemStats(&memAfter)
	cpuUsage := cpuTime() - cpuStart

	// GC pauses
	gcPauses := memAfter.NumGC - memBefore.NumGC

	// Fairness (placeholder)
	fairness := calculateFairness(opTimes, numRoutines)

	return &benchmark.MetricResult{
		MapImpl:      benchmark.LockFree,
		Scenario:     fmt.Sprintf("%d routines × %d ops", numRoutines, opsPerRoutine),
		Throughput:   throughput,
		Latency:      avgLatency,
		Memory:       memAfter.Alloc - memBefore.Alloc,
		CPUUtil:      cpuUsage.Seconds(),
		GCPauses:     int64(gcPauses),
		FairnessData: fairness,
	}
}
