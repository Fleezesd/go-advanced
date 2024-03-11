package GC

import (
	"runtime"
	"testing"
)

// GOGC=off go test -run=Test1000Allocs -trace=trace1.out
// go tool trace trace1.out
func Test1000Allocs(t *testing.T) {
	go func() {
		for {
			i := 123
			reader(&i)
		}
	}()

	for i := 0; i < 1000; i++ {
		ii := i
		i = *reader(&ii)
	}

	runtime.GC()
}

// GOGC=off go test -run=Test10000000000Allocs -trace=trace2.out
// go tool trace trace2.out
func Test10000000000Allocs(t *testing.T) {
	go func() {
		for {
			i := 123
			reader(&i)
		}
	}()

	for i := 0; i < 10000000000; i++ {
		ii := i
		i = *reader(&ii)
	}

	runtime.GC()
}

//go:noinline
func reader(i *int) *int {
	ii := i
	return ii
}
