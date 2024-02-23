package tate

import (
	"context"
	"sync"
)

type Repeater struct {
	wg     sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
}

func NewRepeater() *Repeater {
	ctx, cancel := context.WithCancel(context.Background())
	return &Repeater{
		wg:     sync.WaitGroup{},
		ctx:    ctx,
		cancel: cancel,
	}
}

func (rp *Repeater) Repeat(routine func()) *Repeater {
	rp.wg.Add(1)
	go func() {
		defer rp.wg.Done()
		for {
			select {
			case <-rp.ctx.Done():
				return
			default:
				routine()
			}
		}
	}()
	return rp
}

func (rp *Repeater) Join() {
	rp.cancel()
	rp.wg.Wait()
}
