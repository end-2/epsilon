package epsilon

import (
	"sync/atomic"
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

type Epsilon struct {
	zeroTime time.Time
	pid      uint32
	seq      uint32
}

func New(zeroTime time.Time, pid uint32) *Epsilon {
	return &Epsilon{
		zeroTime: zeroTime,
		pid:      pid,
		seq:      uint32(0),
	}
}

func (e *Epsilon) Now() uint64 {
	return uint64(time.Since(e.zeroTime).Nanoseconds())
}

func (e *Epsilon) Next() uint64 {
	return (e.Now()&TimestampMask)<<TimestampShiftBits +
		uint64(e.pid&ParentsIDMask)<<SequenceNumberBits +
		uint64(atomic.AddUint32(&e.seq, 1)&SequenceNumberMask)
}
