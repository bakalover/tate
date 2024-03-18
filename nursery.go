package tate

type Nursery struct {
	handles []Joinable
}

func NewNursery() *Nursery {
	return &Nursery{}
}

func (n *Nursery) Go(routine func()) *Nursery {
	h := Go(func() { routine() })
	n.handles = append(n.handles, h)
	return n
}

func (n *Nursery) Scope(routine func(sc *Scope)) *Nursery {
	scope := NewScope()
	h := Go(func() {
		defer scope.Join()
		routine(scope)
	})
	n.handles = append(n.handles, h)
	return n
}

func (n *Nursery) Join() {
	for _, h := range n.handles {
		h.Join()
	}
}
