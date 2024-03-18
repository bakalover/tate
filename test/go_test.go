package test

import (
	"testing"

	"github.com/bakalover/tate"
)

func TestGoHandle(t *testing.T) {
	check := false
	h := tate.Go(func() {
		check = true
	})
	h.Join()
	if !check {
		t.Fatal()
	}
}

func TestInsideHandle(t *testing.T) {
	check1 := false
	check2 := false

	h2 := tate.Go(func() { check2 = true })

	h1 := tate.Go(func() {
		check1 = true
		h2.Join()
	})

	h1.Join()

	if !check1 || !check2 {
		t.Fatal()
	}
}
