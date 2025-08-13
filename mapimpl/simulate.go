package mapimpl

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"

	"lf-v-mutex/benchmark"
)

func Simulate(m Impl, numRoutines int, opsPerRoutine int, OpType benchmark.OperationType, mapImpl benchmark.MapType) *benchmark.MetricResult {
	var wg sync.WaitGroup

	start := time.Now()

	// Declaring mem stats and Initializing Current stats
	runtime.GC() // Calling GC to get accurate mem data
	var memBefore, memAfter runtime.MemStats
	runtime.ReadMemStats(&memBefore)

	numOps := numRoutines * opsPerRoutine
	opTimes := make([]int64, 0, numOps)
	cpuStart := CPUTime()

	var opTimesMu sync.Mutex // OpTimes is shared by the Routines

	wg.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go func(id int) {
			defer wg.Done()
			r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(id)))
			read := 5

			if OpType == benchmark.ReadHeavy {
				read = 8
			} else if OpType == benchmark.WriteHeavy {
				read = 2
			}

			localTimes := make([]int64, 0, opsPerRoutine)

			for j := 0; j < opsPerRoutine; j++ {
				t0 := time.Now().UnixNano()

				key := fmt.Sprintf("r%d_op%d", id, j)

				if r.Intn(10) < read {
					m.Load(key)
				} else {
					m.Store(key, j)
				}

				delta := time.Now().UnixNano() - t0
				localTimes = append(localTimes, delta)
			}

			opTimesMu.Lock()
			opTimes = append(opTimes, localTimes...)
			opTimesMu.Unlock()
		}(i)
	}
	wg.Wait()

	totalTime := time.Since(start)
	throughput := float64(numOps) / totalTime.Seconds()
	avgLatency := float64(Sum(opTimes)) / float64(numOps)

	// Memory and CPU stats after
	runtime.GC() // Calling GC to get accurate mem data
	runtime.ReadMemStats(&memAfter)
	cpuUsage := CPUTime() - cpuStart

	return &benchmark.MetricResult{
		OpType:     benchmark.OpTypeMap[OpType],
		MapImpl:    benchmark.MapTypeString[mapImpl],
		Scenario:   fmt.Sprintf("%d routines × %d ops", numRoutines, opsPerRoutine),
		Throughput: throughput,
		Latency:    avgLatency,
		Memory:     memAfter.TotalAlloc - memBefore.TotalAlloc,
		CPUUtil:    cpuUsage.Seconds(),
	}
}
