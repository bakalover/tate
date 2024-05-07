package tate

import "sync/atomic"

type IJoin interface {
	Join()
}

type JoinHandle struct {
	slave   IJoin
	isFired atomic.Bool
}

func NewJoinHandle(j IJoin) *JoinHandle {
	return &JoinHandle{j, atomic.Bool{}}
}

func (h *JoinHandle) Join() {
	if h.isFired.CompareAndSwap(false, true) {
		h.slave.Join()
	} else {
		panic("Double join!")
	}
}
