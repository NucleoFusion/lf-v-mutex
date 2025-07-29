package benchmark

import "fmt"

type MapType int

const (
	LockFree = iota
	RWMut
	CrudeLF
)

var MapTypeString = map[MapType]string{
	LockFree: "LockFree",
	RWMut:    "RWMutex",
	CrudeLF:  "CrudeLF",
}

type MetricResult struct {
	MapImpl      MapType
	Scenario     string
	Throughput   float64
	Latency      float64
	Memory       uint64
	CPUUtil      float64
	GCPauses     int64
	FairnessData []int
}

func (m *MetricResult) Print() {
	fmt.Println("=== Benchmark Results ===")
	fmt.Printf("Map Implementation : %s\n", MapTypeString[m.MapImpl])
	fmt.Printf("Scenario           : %s\n", m.Scenario)
	fmt.Printf("Throughput         : %.2f ops/sec\n", m.Throughput)
	fmt.Printf("Avg Latency        : %.2f µs\n", m.Latency/1000) // assuming latency is in nanoseconds
	fmt.Printf("Memory Used        : %d bytes\n", m.Memory)
	fmt.Printf("CPU Utilization    : %.2fs\n", m.CPUUtil)
	fmt.Printf("GC Pauses          : %d\n", m.GCPauses)
	if len(m.FairnessData) > 0 {
		fmt.Printf("Fairness Data      : %v\n", m.FairnessData)
	}
	fmt.Println("=========================")
}
