package test

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/bakalover/tate"
	"github.com/stretchr/testify/assert"
)

func TestFixScope(t *testing.T) {
	check := false
	tate.FixScope(func(s *tate.Scope) {
		check = true
	})
	assert.True(t, check)
}

func TestFixScopeGo(t *testing.T) {
	check := false
	tate.FixScope(func(s *tate.Scope) {
		s.Go(func() {
			time.Sleep(time.Second * 2)
			check = true
		})
	})
	assert.True(t, check)
}

func TestFixScopeGoGroup(t *testing.T) {
	check1 := false
	check2 := false

	tate.FixScope(func(s *tate.Scope) {
		s.Go(func() {
			time.Sleep(time.Second * 2)
			check1 = true
		})

		s.Go(func() {
			time.Sleep(time.Second * 2)
			check2 = true
		})
	})
	assert.True(t, check1 && check2)
}
func TestDynScope(t *testing.T) {
	check := false
	j := tate.DynScope(func(s *tate.Scope) {
		time.Sleep(2 * time.Second)
		check = true
	})
	assert.True(t, check)
	j.Join()
}

func TestDynScopeGo(t *testing.T) {
	check := atomic.Bool{}
	j := tate.DynScope(func(s *tate.Scope) {
		s.Go(func() {
			time.Sleep(2 * time.Second)
			check.Store(true)
		})
	})
	assert.False(t, check.Load())
	j.Join()
	assert.True(t, check.Load())
}

func TestDynScopeGoGroup(t *testing.T) {
	check1 := atomic.Bool{}
	check2 := atomic.Bool{}
	j := tate.DynScope(func(s *tate.Scope) {
		s.Go(func() {
			time.Sleep(2 * time.Second)
			check1.Store(true)
		})

		s.Go(func() {
			time.Sleep(2 * time.Second)
			check2.Store(true)
		})
	})
	assert.False(t, check1.Load() && check2.Load())
	j.Join()
	assert.True(t, check1.Load() && check2.Load())
}

func TestSubDynScope(t *testing.T) {
	check := false
	j := tate.DynScope(func(s *tate.Scope) {
		innerCheck := false
		tate.FixScope(func(sub *tate.Scope) {
			sub.Go(func() {
				time.Sleep(2 * time.Second)
				innerCheck = true
				check = true
			})
		})
		assert.True(t, innerCheck)
	})
	assert.True(t, check)
	j.Join()
	assert.True(t, check)
}
