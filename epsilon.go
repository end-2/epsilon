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

	MaxPID = uint32((2 << ParentsIDBits) - 1)
	MaxSeq = uint32((2 << SequenceNumberBits) - 1)
)

var (
	ErrPIDExceed      = errors.New("pid value exceeds maximum(2^9-1)")
	ErrSeqNumOverflow = errors.New("sequence number is the maximum! please try again")
)

// Epsilon is a 64-bit ID(composed of timestamp, pid, seq) generator.
type Epsilon struct {
	baseTime time.Time
	pid      uint64
	seq      uint32
	prevTime uint64
	mutex    *sync.Mutex
}

// New function create Epsilon based on base time and pid.
func New(baseTime time.Time, pid uint32) (*Epsilon, error) {
	if pid > MaxPID {
		return nil, ErrPIDExceed
	}

	return &Epsilon{
		baseTime: baseTime,
		pid:      uint64(pid&ParentsIDMask) << SequenceNumberBits,
		seq:      uint32(0),
		prevTime: uint64(0),
		mutex:    &sync.Mutex{},
	}, nil
}

// Now function return epsilon current time.
func (e *Epsilon) Now() uint64 {
	return uint64(time.Since(e.baseTime).Nanoseconds())
}

// Next function return id, if sequential number is full then return ErrSeqNumOverflow.
func (e *Epsilon) Next() (uint64, error) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	timestamp := (e.Now() & TimestampMask) << TimestampShiftBits

	if e.prevTime < timestamp {
		e.prevTime = timestamp
		e.seq = 0
	} else {
		if e.seq > MaxSeq {
			return 0, ErrSeqNumOverflow
		}
		e.seq++
	}

	return timestamp | e.pid | uint64(e.seq&SequenceNumberMask), nil
}
