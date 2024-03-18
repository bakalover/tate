package test

import (
	"testing"
	"time"

	"github.com/bakalover/tate"
)

func TestScope(t *testing.T) {
	check1 := false
	check2 := false
	sc := tate.NewScope()
	sc.
		Go(func() { check1 = true }).
		Go(func() { check2 = true })

	sc.Join()
	if !check1 || !check2 {
		t.Fatal()
	}
}

func TestSubScope(t *testing.T) {
	check1 := false
	check2 := false
	check3 := false
	sc := tate.NewScope()
	sc.Go(func() { check1 = true })
	sc.SubScope(func(sc *tate.Scope) {
		sc.Go(func() { check2 = true })
		sc.Go(func() { check3 = true })
	})
	sc.Join()
	if !check1 || !check2 || !check3 {
		t.Fatal()
	}
}

func TestDoubleSubScope(t *testing.T) {
	check1 := false
	check2 := false
	check3 := false
	sc := tate.NewScope()
	sc.Go(func() { check1 = true })
	sc.SubScope(func(sc *tate.Scope) {
		sc.Go(func() { check2 = true })
		sc.SubScope(func(sc *tate.Scope) {
			check3 = true
		})
	})
	sc.Join()
	if !check1 || !check2 || !check3 {
		t.Fatal()
	}
}

func TestDoubleSleepy(t *testing.T) {
	sc := tate.NewScope()
	sc.Go(func() { time.Sleep(time.Second) })
	sc.SubScope(func(sc *tate.Scope) {
		sc.Go(func() { time.Sleep(time.Second) })
		sc.SubScope(func(sc *tate.Scope) {
			time.Sleep(time.Second)
		})
	})
	sc.Join()
}
