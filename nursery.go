package tate

import "sync"

type Nursery struct {
	wg sync.WaitGroup
}

func (n *Nursery) Add(routine func()) *Nursery {
	n.wg.Add(1)
	go func() {
		routine()
		n.wg.Done()
	}()
	return n
}

func (n *Nursery) Join() {
	n.wg.Wait()
}
