package benchmark

import "fmt"

type MapType int

const (
	LockFree = iota
	RWMut
	Atomic
	PartialLockfree
)

type OperationType int

const (
	Balanced OperationType = iota
	ReadHeavy
	WriteHeavy
)

var OpTypeMap = map[OperationType]string{
	Balanced:   "Balanced",
	ReadHeavy:  "Read Heavy",
	WriteHeavy: "Write Heavy",
}

var MapTypeString = map[MapType]string{
	LockFree:        "LockFree",
	RWMut:           "RWMutex",
	Atomic:          "Atomic",
	PartialLockfree: "Partial Lockfree",
}

type MetricResult struct {
	OpType      string
	MapImpl     string
	Scenario    string
	Throughput  float64
	Latency     float64
	Memory      uint64
	CPUUtil     float64
	SuccessRate float64
}

func (m *MetricResult) Print() {
	fmt.Println("=== Benchmark Results ===")
	fmt.Printf("Operation Type     : %s\n", m.OpType)
	fmt.Printf("Map Implementation : %s\n", m.MapImpl)
	fmt.Printf("Scenario           : %s\n", m.Scenario)
	fmt.Printf("Throughput         : %.2f ops/sec\n", m.Throughput)
	fmt.Printf("Avg Latency        : %.2f µs\n", m.Latency/1000) // assuming latency is in nanoseconds
	fmt.Printf("Memory Used        : %d mb\n", m.Memory/(1024*1024))
	fmt.Printf("CPU Utilization    : %.2fs\n", m.CPUUtil)
	fmt.Println("=========================")
}
