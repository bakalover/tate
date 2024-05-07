package tate

import (
	"sync"
	"sync/atomic"
)

type Gorroutine struct {
	routine func()
	wg      sync.WaitGroup
}

func (g *Gorroutine) Launch() {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		g.routine()
	}()
}

func (g *Gorroutine) Wait() {
	g.wg.Wait()
}

type JoinHandle struct {
	gorr    *Gorroutine
	isFired atomic.Bool
}

func (h *JoinHandle) Join() {
	if !h.isFired.CompareAndSwap(false, true) {
		h.gorr.Wait()
	} else {
		panic("Double join on single goroutine!")
	}
}

// Cancellation??
func Go(routine func()) JoinHandle {
	g := &Gorroutine{routine: routine}
	g.Launch()
	return JoinHandle{g, atomic.Bool{}}
}
