package test

import (
	"sync"
	"testing"
	"time"

	"github.com/bakalover/tate"
	"github.com/stretchr/testify/assert"
)

const IterRepeater = 12325

func TestRepeaterForgetToAdd(t *testing.T) {
	rp := tate.NewRepeater()
	rp.CancelJoin()
}

func TestRepeaterJoins(t *testing.T) {
	rp := tate.NewRepeater()
	rp.CancelJoin()
	rp.CancelJoin()
	rp.CancelJoin()
}

func TestRepeaterGroupStart(t *testing.T) {
	rp := tate.NewRepeater()
	var mutex sync.Mutex
	var counter = 0

	for i := 0; i < IterRepeater; i++ {
		rp.Go(0, func() {
			mutex.Lock()
			counter++
			mutex.Unlock()
		})
	}
	rp.CancelJoin()
	check := counter
	time.Sleep(time.Second)
	assert.Equal(t, check, counter)
	t.Log(counter)
}

func TestRepeaterEachWait(t *testing.T) {
	rp := tate.NewRepeater()
	var mutex sync.Mutex
	var counter = 0

	for i := 0; i < IterRepeater; i++ {
		rp.Go(0, func() {
			mutex.Lock()
			counter++
			mutex.Unlock()
		})
		rp.CancelJoin()
		check := counter
		time.Sleep(time.Microsecond * 3)
		assert.Equal(t, check, counter)
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
			rp.Go(0, func() {
				mutex.Lock()
				counter++
				mutex.Unlock()
			})
		}
		rp.CancelJoin()
		check := counter
		assert.Equal(t, check, counter)
	}

	t.Log(counter)
}
