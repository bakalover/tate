package tate

import "sync"

type Chainer struct {
	toJoin      []Joinable
	routineRefs sync.WaitGroup
}

func (c *Chainer) AddRef() {
	c.routineRefs.Add(1)
}

func (c *Chainer) ReleaseRef() {
	c.routineRefs.Done()
}

func (c *Chainer) Link(js ...Joinable) {
	c.toJoin = append(c.toJoin, js...)
}

func (c *Chainer) Join() {
	c.routineRefs.Wait()
	for _, js := range c.toJoin {
		js.Join()
	}
}

type Nursery struct {
	wg sync.WaitGroup
	cn Chainer
}

func NewNursery(cn *Chainer) *Nursery {
	nr := &Nursery{}
	if cn != nil {
		cn.Link(nr)
	}
	return nr
}

func (n *Nursery) Add(routine func(c *Chainer)) *Nursery {
	n.wg.Add(1)
	n.cn.AddRef()
	go func() {
		defer n.wg.Done()
		routine(&n.cn)
		
		// At this point we have strong knowledge about all subsciptions on Chainer 
		n.cn.ReleaseRef()
	}()
	return n
}

func (n *Nursery) Join() {
	n.cn.Join()
	n.wg.Wait()
}
