package tate

import (
	"sync"
	"testing"
)

const (
	Iter = 452325
)

func Assert(t *testing.T, p bool) {
	if !p {
		t.Fatal("Assert Failed!")
	}
}

func TestJustWork(t *testing.T) {
	var nr Nursery
	var p1, p2 = false, false
	nr.Add(func() { p1 = true }).Add(func() { p2 = true }).Join()
	Assert(t, p1)
	Assert(t, p2)
}

func TestForgetToAdd(t *testing.T) {
	var nr Nursery
	nr.Join()
}

func TestJoins(t *testing.T) {
	var nr Nursery
	nr.Join()
	nr.Join()
	nr.Join()
	nr.Join()
}

func TestSeveralJoins(t *testing.T) {
	var nr Nursery
	var p1, p2, p3 = false, false, false

	nr.Add(func() { p1 = true })
	nr.Join()
	Assert(t, p1)

	nr.Add(func() { p2 = true })
	nr.Join()
	Assert(t, p2)

	nr.Add(func() { p3 = true })
	nr.Join()
	Assert(t, p3)
}

func TestGroupStart(t *testing.T) {
	var nr Nursery
	var mutex sync.Mutex
	var counter = 0

	for i := 0; i < Iter; i++ {
		nr.Add(func() {
			mutex.Lock()
			counter++
			mutex.Unlock()
		})
	}
	nr.Join()
	Assert(t, counter == Iter)
}

func TestEachWait(t *testing.T) {
	var nr Nursery
	var mutex sync.Mutex
	var counter = 0

	for i := 0; i < Iter; i++ {
		nr.Add(func() {
			mutex.Lock()
			counter++
			mutex.Unlock()
		})
		nr.Join()
	}

	Assert(t, counter == Iter)
}

func TestBatchWait(t *testing.T) {
	var nr Nursery
	var mutex sync.Mutex
	var counter = 0
	var kBatch = 25

	for i := 0; i < kBatch; i++ {
		for i := 0; i < Iter/kBatch; i++ {
			nr.Add(func() {
				mutex.Lock()
				counter++
				mutex.Unlock()
			})
		}
		nr.Join()
	}

	Assert(t, counter == Iter)
}
