package epsilon

import (
	"sync"
	"testing"
	"time"
)

var TimeBoundary = time.Nanosecond

func TestNext(t *testing.T) {
	testCases := 100000
	idList := make([]uint64, testCases)

	got := 0
	expected := 0

	e := New(time.Now(), 0)

	for i := 0; i < testCases; {
		time.Sleep(TimeBoundary)
		if id, err := e.Next(); err != nil {
			continue
		} else {
			idList[i] = id
			i++
		}
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

func TestNext1000Thread(t *testing.T) {
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
			for i := 0; i < testCases/thread; {
				time.Sleep(TimeBoundary)
				if id, err := e.Next(); err != nil {
					continue
				} else {
					idList[testCases/thread*tt+i] = id
					i++
				}
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

func TestNext10000Thread(t *testing.T) {
	testCases := 100000
	thread := 10000
	idList := make([]uint64, testCases)
	wg := sync.WaitGroup{}

	got := 0
	expected := 0

	e := New(time.Now(), 0)

	for tt := 0; tt < thread; tt++ {
		go func(tt int) {
			wg.Add(1)
			defer wg.Done()
			for i := 0; i < testCases/thread; {
				time.Sleep(TimeBoundary)
				if id, err := e.Next(); err != nil {
					continue
				} else {
					idList[testCases/thread*tt+i] = id
					i++
				}
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

func TestNextMultiParents10000Thread(t *testing.T) {
	testCases := 100000
	thread := 10000
	maxPid := uint32(512)
	idList := make([]uint64, testCases)
	wg := sync.WaitGroup{}

	got := 0
	expected := 0

	zt := time.Now()
	es := make([]*Epsilon, maxPid)
	for p := uint32(0); p < maxPid; p++ {
		es[p] = New(zt, p)
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
					idList[testCases/thread*tt+i] = id
					i++
				}
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

// TestNextPrint is to verify the generated id as a bit.
//func TestNextPrint(t *testing.T) {
//	testCases := 1000
//	thread := 10
//	maxPid := uint32(8)
//	idList := make([]uint64, testCases)
//	wg := sync.WaitGroup{}
//
//	got := 0
//	expected := 0
//
//	zt := time.Now()
//	es := make([]*Epsilon, maxPid)
//	for p := uint32(0); p < maxPid; p++ {
//		time.Sleep(time.Microsecond)
//		es[p] = New(zt, p)
//	}
//	for tt := 0; tt < thread; tt++ {
//		go func(tt int) {
//			wg.Add(1)
//			defer wg.Done()
//
//			e := es[tt%int(maxPid)]
//			for i := 0; i < testCases/thread; {
//				time.Sleep(TimeBoundary)
//				if id, err := e.Next(); err != nil {
//					continue
//				} else {
//					idList[testCases/thread*tt+i] = id
//					i++
//				}
//			}
//		}(tt)
//	}
//	wg.Wait()
//
//	sort.Slice(idList, func(i, j int) bool {
//		if idList[i] < idList[j] {
//			return true
//		}
//		return false
//	})
//	for i := 0; i < testCases; i++ {
//		fmt.Printf("%b\n", idList[i])
//	}
//
//	for i := 0; i < testCases; i++ {
//		for j := i + 1; j < testCases; j++ {
//			if idList[i] == idList[j] {
//				got++
//			}
//		}
//	}
//	if got != expected {
//		t.Errorf("expected %d, got %d", expected, got)
//	}
//}
