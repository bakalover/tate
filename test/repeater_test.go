package test

import (
	"sync"
	"testing"
	"time"

	"github.com/bakalover/tate"
)

const IterRepeater = 123125

func TestRepeaterForgetToAdd(t *testing.T) {
	rp := tate.NewRepeater()
	rp.Join()
}

func TestRepeaterJoins(t *testing.T) {
	rp := tate.NewRepeater()
	rp.Join()
	rp.Join()
	rp.Join()
	rp.Join()
}

func TestRepeaterGroupStart(t *testing.T) {
	rp := tate.NewRepeater()
	var mutex sync.Mutex
	var counter = 0

	for i := 0; i < IterRepeater; i++ {
		rp.Repeat(func() {
			mutex.Lock()
			counter++
			mutex.Unlock()
		})
	}
	rp.Join()
	check := counter
	time.Sleep(time.Second)
	if check != counter {
		t.Fatal()
	}
	t.Log(counter)
}

func TestRepeaterEachWait(t *testing.T) {
	rp := tate.NewRepeater()
	var mutex sync.Mutex
	var counter = 0

	for i := 0; i < IterRepeater; i++ {
		rp.Repeat(func() {
			mutex.Lock()
			counter++
			mutex.Unlock()
		})
		rp.Join()
		check := counter
		time.Sleep(time.Microsecond * 3)
		if check != counter {
			t.Fatal()
		}
	}
	t.Log(counter)
}

func TestRepeaterBatchWait(t *testing.T) {
	rp := tate.NewRepeater()
	var mutex sync.Mutex
	var counter = 0
	var kBatch = 25

	for i := 0; i < kBatch; i++ {
		for i := 0; i < IterRepeater/kBatch; i++ {
			rp.Repeat(func() {
				mutex.Lock()
				counter++
				mutex.Unlock()
			})
		}
		rp.Join()
		check := counter
		time.Sleep(time.Second)
		if check != counter {
			t.Fatal()
		}
	}

	t.Log(counter)
}
