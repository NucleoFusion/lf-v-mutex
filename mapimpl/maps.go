package mapimpl

import (
	"syscall"
	"time"
)

func sum(arr []int64) int64 {
	var s int64
	for _, v := range arr {
		s += v
	}
	return s
}

// Placeholder: More advanced fairness metrics can be added
func calculateFairness(opTimes []int64, numRoutines int) []int {
	// Count how many operations per routine (if you track that separately)
	// Here we simply return percentiles or dummy data
	return []int{25, 50, 75} // Replace with actual fairness calculations
}

// Approximate CPU time for current process
func cpuTime() time.Duration {
	var ru syscall.Rusage
	syscall.Getrusage(syscall.RUSAGE_SELF, &ru)
	return time.Duration(ru.Utime.Sec)*time.Second +
		time.Duration(ru.Utime.Usec)*time.Microsecond
}
