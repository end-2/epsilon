package epsilon

import (
	"sync"
	"testing"
	"time"
)

var TimeBoundary = 10 * time.Microsecond

func TestNext(t *testing.T) {
	testCases := 100000
	idList := make([]uint64, testCases)

	got := 0
	expected := 0

	e := New(time.Now(), 0)

	for i := 0; i < testCases; i++ {
		time.Sleep(TimeBoundary)
		idList[i] = e.Next()
	}

	for i := 0; i < testCases; i++ {
		for j := i + 1; j < testCases; j++ {
			if idList[i] == idList[j] {
				got++
			}
		}
	}
	if got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}
}

func TestNextMultiThread(t *testing.T) {
	testCases := 100000
	thread := 1000
	idList := make([]uint64, testCases)
	wg := sync.WaitGroup{}

	got := 0
	expected := 0

	e := New(time.Now(), 0)

	for tt := 0; tt < thread; tt++ {
		go func(tt int) {
			wg.Add(1)
			defer wg.Done()
			for i := 0; i < testCases/thread; i++ {
				time.Sleep(TimeBoundary)
				idList[testCases/thread*tt+i] = e.Next()
			}
		}(tt)
	}
	wg.Wait()

	for i := 0; i < testCases; i++ {
		for j := i + 1; j < testCases; j++ {
			if idList[i] == idList[j] {
				got++
			}
		}
	}
	if got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}
}
