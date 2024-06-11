package tate

import (
	"sync"
)

type Gorroutine struct {
	routine func(...any)
	args    []any
	wg      sync.WaitGroup
}

func (g *Gorroutine) Launch() {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		g.routine(g.args...)
	}()
}

func (g *Gorroutine) Join() {
	g.wg.Wait()
}

func Go(routine func(...any), args ...any) *JoinHandle {
	g := &Gorroutine{routine: routine, args: args}
	g.Launch()
	return NewJoinHandle(g)
}
