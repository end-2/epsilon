package epsilon

import (
	"errors"
	"sync"
	"time"
)

const (
	TimestampBits      = 45
	ParentsIDBits      = 9
	SequenceNumberBits = 10

	TimestampCutOffBits = 16
	TimestampShiftBits  = 64 - TimestampCutOffBits - TimestampBits
	TimestampMask       = ^(uint64(1)<<TimestampCutOffBits - 1)
	ParentsIDMask       = uint32(1)<<ParentsIDBits - 1
	SequenceNumberMask  = uint32(1)<<SequenceNumberBits - 1
)

var (
	ErrSeqNumOverflow = errors.New("sequence number is the maximum! please try again")
)

// Epsilon is a 64-bit ID(composed of timestamp, pid, seq) generator.
type Epsilon struct {
	zeroTime time.Time
	pid      uint64
	seq      uint32
	pre      uint64
	mutex    sync.Mutex
}

// New function create Epsilon based on base time and pid.
func New(zeroTime time.Time, pid uint32) *Epsilon {
	return &Epsilon{
		zeroTime: zeroTime,
		pid:      uint64(pid&ParentsIDMask) << SequenceNumberBits,
		seq:      uint32(0),
		pre:      uint64(0),
		mutex:    sync.Mutex{},
	}
}

// Now function return epsilon current time.
func (e *Epsilon) Now() uint64 {
	return uint64(time.Since(e.zeroTime).Nanoseconds())
}

// Next function return id, if sequential number is full then return ErrSeqNumOverflow.
func (e *Epsilon) Next() (uint64, error) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	timestamp := (e.Now() & TimestampMask) << TimestampShiftBits
	e.seq++
	if e.pre < timestamp {
		e.pre = timestamp
		e.seq = 1
	}
	if e.seq == 0 {
		return 0, ErrSeqNumOverflow
	}
	return timestamp + e.pid + uint64(e.seq&SequenceNumberMask), nil
}
