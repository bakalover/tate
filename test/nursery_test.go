package test

import (
	"sync"
	"testing"

	"github.com/bakalover/tate"
)

const IterNursery = 123225

func TestNurseryJustWork(t *testing.T) {
	nr := tate.NewNursery()
	var check1, check2 = false, false

	nr.Go(func() { check1 = true })
	nr.Go(func() { check2 = true })

	nr.Join()

	if !check1 || !check2 {
		t.Fatal()
	}
}

func TestNurseryForgetToAdd(t *testing.T) {
	nr := tate.NewNursery()
	nr.Join()
}

func TestNurseryManyJoins(t *testing.T) {
	nr := tate.NewNursery()
	nr.Join()
	nr.Join()
	nr.Join()
	nr.Join()
}

func TestNurseryReUse(t *testing.T) {
	nr := tate.NewNursery()
	var check1, check2, check3 = false, false, false

	nr.Go(func() { check1 = true })
	nr.Join()
	if !check1 {
		t.Fatal()
	}

	nr.Go(func() { check2 = true })
	nr.Join()
	if !check2 {
		t.Fatal()
	}

	nr.Go(func() { check3 = true })
	nr.Join()
	if !check3 {
		t.Fatal()
	}
}

func TestNurseryGroup(t *testing.T) {
	nr := tate.NewNursery()
	var mutex sync.Mutex
	var counter = 0

	for i := 0; i < IterNursery; i++ {
		nr.Go(func() {
			mutex.Lock()
			defer mutex.Unlock()
			counter++
		})
	}

	nr.Join()
	if !(counter == IterNursery) {
		t.Fatal()
	}
}

func TestNurseryBatchWait(t *testing.T) {
	nr := tate.NewNursery()
	var mutex sync.Mutex
	var counter = 0
	var kBatch = 25

	for i := 0; i < kBatch; i++ {
		for i := 0; i < IterNursery/kBatch; i++ {
			nr.Go(func() {
				mutex.Lock()
				defer mutex.Unlock()
				counter++
			})
		}
		nr.Join()
	}

	if !(counter == IterNursery) {
		t.Fatal()
	}
}

func TestNurseryScope(t *testing.T) {
	nr := tate.NewNursery()
	check := false
	nr.Scope(func(sc *tate.Scope) {
		check = true
	})
	nr.Join()
	if !check {
		t.Fatal()
	}
}

func TestNurseryGoAndScope(t *testing.T) {
	nr := tate.NewNursery()
	check1, check2 := false, false

	nr.Go(func() {
		check1 = true
	})

	nr.Scope(func(sc *tate.Scope) {
		check2 = true
	})
	nr.Join()
	if !check1 || !check2 {
		t.Fatal()
	}
}

func TestNurseryTelescope(t *testing.T) {
	nr := tate.NewNursery()
	check := false

	nr.Scope(func(sc1 *tate.Scope) {
		sc1.SubScope(func(sc2 *tate.Scope) {
			sc2.SubScope(func(sc3 *tate.Scope) {
				check = true
			})
		})
	})

	nr.Join()
	if !check {
		t.Fatal()
	}
}

func TestConcurrentTree(t *testing.T) {
	nr := tate.NewNursery()
	var mutex sync.Mutex
	var counter = 0

	nr.Scope(func(sc *tate.Scope) {
		sc.Go(func() {

			for i := 0; i < IterNursery; i++ {
				mutex.Lock()
				counter += 1
				mutex.Unlock()
			}
		})
		sc.Go(func() {
			for i := 0; i < IterNursery; i++ {
				mutex.Lock()
				counter ++
				mutex.Unlock()
			}
		})
	})

	nr.Scope(func(sc *tate.Scope) {
		sc.Go(func() {
			for i := 0; i < IterNursery; i++ {
				mutex.Lock()
				counter ++
				mutex.Unlock()
			}
		})
		sc.SubScope(func(sc *tate.Scope) {
			for i := 0; i < IterNursery; i++ {
				mutex.Lock()
				counter ++
				mutex.Unlock()
			}
		})
	})

	nr.Join()
	if counter != IterNursery*4 {
		t.Fatal()
	}
}
