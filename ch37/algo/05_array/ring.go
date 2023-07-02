package _05_array

import (
	"errors"
	log "github.com/golang/glog"
)

type Proto struct {
	Ver  int32  `protobuf:"varint,1,opt,name=ver,proto3" json:"ver,omitempty"`
	Op   int32  `protobuf:"varint,2,opt,name=op,proto3" json:"op,omitempty"`
	Seq  int32  `protobuf:"varint,3,opt,name=seq,proto3" json:"seq,omitempty"`
	Body []byte `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

// Ring ring proto buffer.
type Ring struct {
	// read
	rp   uint64
	num  uint64
	mask uint64
	// TODO split cacheline, many cpu cache line size is 64
	// pad [40]byte
	// write
	wp   uint64
	data []Proto
}

// NewRing new a ring buffer.
func NewRing(num int) *Ring {
	r := new(Ring)
	r.init(uint64(num))
	return r
}

// Init init ring.
func (r *Ring) Init(num int) {
	r.init(uint64(num))
}

func (r *Ring) init(num uint64) {
	// 2^N
	if num&(num-1) != 0 {
		for num&(num-1) != 0 {
			num &= num - 1
		}
		num <<= 1
	}
	r.data = make([]Proto, num)
	r.num = num
	r.mask = r.num - 1
}

// Get get a proto from ring.
func (r *Ring) Get() (proto *Proto, err error) {
	if r.rp == r.wp {
		return nil, errors.New("ring buffer empty")
	}
	proto = &r.data[r.rp&r.mask]
	return
}

// GetAdv incr read index.
func (r *Ring) GetAdv() {
	r.rp++
	if true {
		log.Infof("ring rp: %d, idx: %d", r.rp, r.rp&r.mask)
	}
}

// Set get a proto to write.
func (r *Ring) Set() (proto *Proto, err error) {
	if r.wp-r.rp >= r.num {
		return nil, errors.New("ring buffer full")
	}

	proto = &r.data[r.wp&r.mask]
	return
}

// SetAdv incr write index.
func (r *Ring) SetAdv() {
	r.wp++
	if true {
		log.Infof("ring wp: %d, idx: %d", r.wp, r.wp&r.mask)
	}
}

// Reset reset ring.
func (r *Ring) Reset() {
	r.rp = 0
	r.wp = 0
	// prevent pad compiler optimization
	// r.pad = [40]byte{}
}
