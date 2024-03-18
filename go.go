package tate

import "sync"

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
	gorr *Gorroutine
}

func (h JoinHandle) Join() {
	h.gorr.Wait()
}

func Go(routine func()) JoinHandle {
	g := &Gorroutine{routine: routine}
	g.Launch()
	return JoinHandle{g}
}
