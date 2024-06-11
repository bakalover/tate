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
		rp.Go(0, func(args ...any) {
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
