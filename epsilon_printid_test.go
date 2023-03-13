package epsilon

import (
	"sort"
	"sync"
	"testing"
	"time"
)

//TestNextPrint is to verify the generated id as a bit.
func TestNextPrint(t *testing.T) {
	testCases := 1000
	thread := 100
	maxPid := uint32(16)

	idBuffer := make(chan uint64, 1048576)

	wg := &sync.WaitGroup{}

	got := false

	zt := time.Now()
	es := make([]*Epsilon, maxPid)
	for p := uint32(0); p < maxPid; p++ {
		var err error
		es[p], err = New(zt, p)
		if err != nil {
			t.Fatal(err)
		}
	}

	for tt := 0; tt < thread; tt++ {
		go func(tt int) {
			wg.Add(1)
			defer wg.Done()

			e := es[tt%int(maxPid)]
			for i := 0; i < testCases/thread; {
				time.Sleep(TimeBoundary)
				if id, err := e.Next(); err != nil {
					continue
				} else {
					if id == uint64(0) {
						t.Error("generated 0 value")
					}
					idBuffer <- id
					i++
				}
			}
		}(tt)
	}
	wg.Wait()

	allList := make([]uint64, testCases)
	for i := 0; i < testCases; i++ {
		allList[i] = <-idBuffer
	}

	sort.Slice(allList, func(i, j int) bool {
		if allList[i] < allList[j] {
			return true
		}
		return false
	})

	for i := 0; i < testCases; i++ {
		for j := i + 1; j < testCases; j++ {
			if allList[i] == allList[j] {
				got = true
			}
		}
	}
	if got {
		t.Error("The same value has occurred.")
	}

	for i := 0; i < testCases; i++ {
		t.Logf("%.64b\n", allList[i])
	}
}
