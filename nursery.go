package tate

import "sync"

type Linker struct {
	toJoin          []Joinable
	routinesLinks sync.WaitGroup
}

func (c *Linker) AddRef() {
	c.routinesLinks.Add(1)
}

func (c *Linker) ReleaseRef() {
	c.routinesLinks.Done()
}

func (c *Linker) Link(js ...Joinable) {
	c.toJoin = append(c.toJoin, js...)
}

func (c *Linker) Join() {
	c.routinesLinks.Wait()
	for _, js := range c.toJoin {
		js.Join()
	}
}

type Nursery struct {
	wg sync.WaitGroup
	cn Linker
}

func NewNursery(cn *Linker) *Nursery {
	nr := &Nursery{}
	if cn != nil {
		cn.Link(nr)
	}
	return nr
}

func (n *Nursery) Add(routine func(c *Linker)) *Nursery {
	n.wg.Add(1)
	n.cn.AddRef()
	go func() {
		defer n.wg.Done()
		routine(&n.cn)

		// At this point we have strong knowledge about all subsciptions on Linker
		n.cn.ReleaseRef()
	}()
	return n
}

func (n *Nursery) Join() {
	n.cn.Join()
	n.wg.Wait()
}
