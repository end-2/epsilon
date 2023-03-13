package epsilon

import (
	"sync"
	"testing"
	"time"
)

var TimeBoundary = time.Nanosecond

func TestNext(t *testing.T) {
	testCases := 10000
	res := make(map[uint64]struct{})

	e, err := New(time.Now(), 0)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < testCases; i++ {
		time.Sleep(TimeBoundary)
		if id, err := e.Next(); err != nil {
			t.Error(err)
			continue
		} else {
			if _, ok := res[id]; ok {
				t.Error("got non-unique id")
			}
			res[id] = struct{}{}
		}
	}
}

func TestNext1000Thread(t *testing.T) {
	testCases := 10000
	thread := 1000
	idBuffer := make(chan uint64, 1048576)
	res := make(map[uint64]struct{})
	wg := &sync.WaitGroup{}

	e, err := New(time.Now(), 0)
	if err != nil {
		t.Fatal(err)
	}

	for tt := 0; tt < thread; tt++ {
		go func(tt int) {
			wg.Add(1)
			defer wg.Done()
			for i := 0; i < testCases/thread; {
				time.Sleep(TimeBoundary)
				if id, err := e.Next(); err != nil {
					continue
				} else {
					idBuffer <- id
					i++
				}
			}
		}(tt)
	}
	wg.Wait()

	for i := 0; i < testCases; i++ {
		id := <-idBuffer
		if _, ok := res[id]; ok {
			t.Error("got non-unique id")
		}
		res[id] = struct{}{}
	}
}

func TestNextMultiParents1000Thread(t *testing.T) {
	testCases := 10000
	thread := 1000
	maxPid := uint32(511)
	idBuffer := make(chan uint64, 1048576)
	res := make(map[uint64]struct{})
	wg := &sync.WaitGroup{}

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
					idBuffer <- id
					i++
				}
			}
		}(tt)
	}
	wg.Wait()

	for i := 0; i < testCases; i++ {
		id := <-idBuffer
		if _, ok := res[id]; ok {
			t.Error("got non-unique id")
		}
		res[id] = struct{}{}
	}
}

func TestNextMultiParents5000Thread(t *testing.T) {
	testCases := 50000
	thread := 5000
	maxPid := uint32(511)
	idBuffer := make(chan uint64, 1048576)
	res := make(map[uint64]struct{})
	wg := &sync.WaitGroup{}

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
					idBuffer <- id
					i++
				}
			}
		}(tt)
	}
	wg.Wait()

	for i := 0; i < testCases; i++ {
		id := <-idBuffer
		if _, ok := res[id]; ok {
			t.Error("got non-unique id")
		}
		res[id] = struct{}{}
	}
}
