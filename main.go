package main

import (
	"encoding/json"
	"fmt"
	"os"

	"lf-v-mutex/benchmark"
)

type Results struct {
	LockFree        []Result `json:"lockFree"`
	RWMutex         []Result `json:"rwMutex"`
	PartialLockfree []Result `json:"partialLockfree"`
}

type Result struct {
	ReadHeavy  *benchmark.MetricResult `json:"readHeavy"`
	WriteHeavy *benchmark.MetricResult `json:"writeHeavy"`
	Balanced   *benchmark.MetricResult `json:"balanced"`
}

func main() {
	fmt.Println("Benchmarking...")
	results := Run()
	fmt.Println("Results Obtained")

	fmt.Println("Converting to JSON...")

	// Lock-Free
	f, _ := os.Create("./results/lockfree.json")
	data, _ := json.Marshal(&results.LockFree)
	f.Write(data)
	f.Close()

	// RWMutex
	f, _ = os.Create("./results/rwmutex.json")
	data, _ = json.Marshal(&results.RWMutex)
	f.Write(data)
	f.Close()

	// PartialLockfree
	f, _ = os.Create("./results/partiallockfree.json")
	data, _ = json.Marshal(&results.PartialLockfree)
	f.Write(data)
	f.Close()

	fmt.Println("Converted to JSON")
	fmt.Println("Successfully Completed Benchmark")
}
