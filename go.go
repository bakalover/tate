package tate

import (
	"sync"
)

type goroutine struct {
	routine func(args ...any)
	args    []any
	wg      sync.WaitGroup
}

func (g *goroutine) Launch() {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		g.routine(g.args...)
	}()
}

func (g *goroutine) Join() {
	g.wg.Wait()
}

func Go(routine func(args ...any), args ...any) *JoinHandle {
	g := &goroutine{routine: routine, args: args}
	g.Launch()
	return NewJoinHandle(g)
}
