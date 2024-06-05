package test

import (
	"testing"

	"github.com/bakalover/tate"
	"github.com/stretchr/testify/assert"
)

func TestGoHandle(t *testing.T) {
	check := false
	h := tate.Go(func(...any) {
		check = true
	})
	h.Join()
	assert.True(t, check)
}

func TestShouldPanic(t *testing.T) {
	h := tate.Go(func(...any) {})
	h.Join()
	assert.Panics(t, func() { h.Join() })
	assert.Panics(t, func() { h.Join() })
}

func InnerHandleTest(t *testing.T) {
	check1 := false
	check2 := false

	h2 := tate.Go(func(...any) { check2 = true })

	h1 := tate.Go(func(...any) {
		check1 = true
		h2.Join()
	})

	h1.Join()
	assert.True(t, check1 && check2)
}
