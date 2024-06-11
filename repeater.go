package tate

import (
	"context"
	"time"
)

type Repeater struct {
	handles []IJoin
	ctx     context.Context
	cancel  context.CancelFunc
}

func NewRepeater() *Repeater {
	ctx, cancel := context.WithCancel(context.Background())
	return &Repeater{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (rp *Repeater) Go(d time.Duration, routine func(args ...any), args ...any) *Repeater {
	h := Go(func(args ...any) {
		for {
			select {
			case <-rp.ctx.Done():
				return
			default:
				routine(args...)
			}
			time.Sleep(d)
		}
	})
	rp.handles = append(rp.handles, h)
	return rp
}

func (rp *Repeater) CancelJoin() {
	rp.cancel()
	for _, h := range rp.handles {
		h.Join()
	}
}
