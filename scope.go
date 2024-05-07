package tate

type Scope struct {
	handles []IJoin
}

func FixScope(f func(s *Scope)) {
	sc := &Scope{}
	f(sc)     // Synchronous spawn
	sc.Join() // Await IJoins spawned inside Scope
}

func DynScope(f func(s *Scope)) *JoinHandle {
	sc := &Scope{}
	f(sc)
	return NewJoinHandle(sc) // Join where we need
}

func (s *Scope) Go(routine func()) *Scope {
	h := Go(func() { routine() })
	s.handles = append(s.handles, h)
	return s
}


func (s *Scope) Join() {
	for _, h := range s.handles {
		h.Join()
	}
}
