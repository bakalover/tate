package test

import (
	"sync"
	"testing"

	"github.com/bakalover/tate"
)

const IterNursery = 123125

func TestNurseryJustWork(t *testing.T) {
	var nr tate.Nursery
	var p1, p2 = false, false
	nr.Add(func() { p1 = true }).Add(func() { p2 = true }).Join()
	if !p1 || !p2 {
		t.Fatal()
	}
}

func TestNurseryForgetToAdd(t *testing.T) {
	var nr tate.Nursery
	nr.Join()
}

func TestNurseryJoins(t *testing.T) {
	var nr tate.Nursery
	nr.Join()
	nr.Join()
	nr.Join()
	nr.Join()
}

func TestNurserySeveralJoins(t *testing.T) {
	var nr tate.Nursery
	var p1, p2, p3 = false, false, false

	nr.Add(func() { p1 = true })
	nr.Join()
	if !p1 {
		t.Fatal()
	}

	nr.Add(func() { p2 = true })
	nr.Join()
	if !p2 {
		t.Fatal()
	}

	nr.Add(func() { p3 = true })
	nr.Join()
	if !p3 {
		t.Fatal()
	}
}

func TestNurseryGroupStart(t *testing.T) {
	var nr tate.Nursery
	var mutex sync.Mutex
	var counter = 0

	for i := 0; i < IterNursery; i++ {
		nr.Add(func() {
			mutex.Lock()
			counter++
			mutex.Unlock()
		})
	}
	nr.Join()
	if !(counter == IterNursery) {
		t.Fatal()
	}
}

func TestNurseryEachWait(t *testing.T) {
	var nr tate.Nursery
	var mutex sync.Mutex
	var counter = 0

	for i := 0; i < IterNursery; i++ {
		nr.Add(func() {
			mutex.Lock()
			counter++
			mutex.Unlock()
		})
		nr.Join()
	}

	if !(counter == IterNursery) {
		t.Fatal()
	}
}

func TestNurseryBatchWait(t *testing.T) {
	var nr tate.Nursery
	var mutex sync.Mutex
	var counter = 0
	var kBatch = 25

	for i := 0; i < kBatch; i++ {
		for i := 0; i < IterNursery/kBatch; i++ {
			nr.Add(func() {
				mutex.Lock()
				counter++
				mutex.Unlock()
			})
		}
		nr.Join()
	}

	if !(counter == IterNursery) {
		t.Fatal()
	}
}
