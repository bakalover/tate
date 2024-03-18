package tate

type Scope struct {
	handles []Joinable
}

func NewScope() *Scope {
	return &Scope{}
}

func (sc *Scope) Join() {
	for _, h := range sc.handles {
		h.Join()
	}
}

func (sc *Scope) Go(routine func()) *Scope {
	h := Go(func() { routine() })
	sc.handles = append(sc.handles, h)
	return sc
}

func (sc *Scope) SubScope(routine func(sub *Scope)) {
	subScope := NewScope()
	h := Go(func() {
		defer subScope.Join()
		routine(subScope)
	})
	sc.handles = append(sc.handles, h)
}
