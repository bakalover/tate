package test

import (
	"testing"

	"github.com/bakalover/tate"
	"github.com/stretchr/testify/assert"
)

func TestGoHandle(t *testing.T) {
	check := false
	h := tate.Go(func(args ...any) {
		check = true
	})
	h.Join()
	assert.True(t, check)
}

func TestShouldPanic(t *testing.T) {
	h := tate.Go(func(args ...any) {})
	h.Join()
	assert.Panics(t, func() { h.Join() })
	assert.Panics(t, func() { h.Join() })
}

func TestInnerHandle(t *testing.T) {
	check1 := false
	check2 := false

	h2 := tate.Go(func(args ...any) { check2 = true })

	h1 := tate.Go(func(args ...any) {
		check1 = true
		h2.Join()
	})

	h1.Join()
	assert.True(t, check1 && check2)
}

func FuzzArgs(f *testing.F) {
	f.Add(912878)
	f.Fuzz(func(t *testing.T, a int) {
		tate.Go(func(args ...any) {
			assert.Equal(t, a, args[0].(int))
		}, a)
	})
}
