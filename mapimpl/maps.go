package mapimpl

import (
	"syscall"
	"time"
)

type Impl interface {
	Load(string) (int, bool)
	Store(string, int)
}

func Sum(arr []int64) int64 {
	var s int64
	for _, v := range arr {
		s += v
	}
	return s
}

// Gets CPU time consumed
func CPUTime() time.Duration {
	var ru syscall.Rusage
	syscall.Getrusage(syscall.RUSAGE_SELF, &ru)
	return time.Duration(ru.Utime.Sec)*time.Second +
		time.Duration(ru.Utime.Usec)*time.Microsecond
}
