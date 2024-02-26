package test

import (
	"sync"
	"testing"
	"time"

	"github.com/bakalover/tate"
)

const IterNursery = 123125

func TestNurseryJustWork(t *testing.T) {
	nr := tate.NewNursery(nil)
	var p1, p2 = false, false
	nr.Add(func(c *tate.Linker) { p1 = true }).Add(func(c *tate.Linker) { p2 = true }).Join()
	if !p1 || !p2 {
		t.Fatal()
	}
}

func TestNurseryForgetToAdd(t *testing.T) {
	nr := tate.NewNursery(nil)
	nr.Join()
}

func TestNurseryJoins(t *testing.T) {
	nr := tate.NewNursery(nil)
	nr.Join()
	nr.Join()
	nr.Join()
	nr.Join()
}

func TestNurserySeveralJoins(t *testing.T) {
	nr := tate.NewNursery(nil)
	var p1, p2, p3 = false, false, false

	nr.Add(func(c *tate.Linker) { p1 = true })
	nr.Join()
	if !p1 {
		t.Fatal()
	}

	nr.Add(func(c *tate.Linker) { p2 = true })
	nr.Join()
	if !p2 {
		t.Fatal()
	}

	nr.Add(func(c *tate.Linker) { p3 = true })
	nr.Join()
	if !p3 {
		t.Fatal()
	}
}

func TestNurseryGroupStart(t *testing.T) {
	nr := tate.NewNursery(nil)
	var mutex sync.Mutex
	var counter = 0

	for i := 0; i < IterNursery; i++ {
		nr.Add(func(c *tate.Linker) {
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
	nr := tate.NewNursery(nil)
	var mutex sync.Mutex
	var counter = 0

	for i := 0; i < IterNursery; i++ {
		nr.Add(func(c *tate.Linker) {
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
	nr := tate.NewNursery(nil)
	var mutex sync.Mutex
	var counter = 0
	var kBatch = 25

	for i := 0; i < kBatch; i++ {
		for i := 0; i < IterNursery/kBatch; i++ {
			nr.Add(func(c *tate.Linker) {
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

func TestChain(t *testing.T) {
	nr := tate.NewNursery(nil)
	p := false
	nr.Add(func(c *tate.Linker) {
		nrInner := tate.NewNursery(c)
		nrInner.Add(func(c *tate.Linker) { p = true })
	})
	nr.Join()
	if !p {
		t.Fatal()
	}
}

func TestLongSubsription(t *testing.T) {
	nr := tate.NewNursery(nil)
	p := false
	nr.Add(func(c *tate.Linker) {
		nrInner := tate.NewNursery(c)
		nrInner.Add(func(c *tate.Linker) {
			time.Sleep(5 * time.Second)
			p = true
		})
	})
	nr.Join()
	if !p {
		t.Fatal()
	}
}

func TestStrangeApi(t *testing.T) {
	nr := tate.NewNursery(nil)
	p := false
	nr.Add(func(c *tate.Linker) {
		nrInner := tate.NewNursery(c)
		nrInner.Join()
		nrInner.Add(func(c *tate.Linker) {
			p = true
		})
		nrInner.Join()
		nrInner.Join()
		nrInner.Join()
	})
	nr.Join()
	if !p {
		t.Fatal()
	}
}

func TestTelescope(t *testing.T) {
	nr1 := tate.NewNursery(nil)
	p := false
	nr1.Add(func(c1 *tate.Linker) {
		nr2 := tate.NewNursery(c1)
		nr2.Add(func(c2 *tate.Linker) {
			nr3 := tate.NewNursery(c2)
			nr3.Add(func(c3 *tate.Linker) {
				nr4 := tate.NewNursery(c3)
				nr4.Add(func(c4 *tate.Linker) {
					nr5 := tate.NewNursery(c4)
					nr5.Add(func(c5 *tate.Linker) {
						nr6 := tate.NewNursery(c5)
						nr6.Add(func(c6 *tate.Linker) {
							p = true
						})
					})
				})
			})
		})
	})
	nr1.Join()
	if !p {
		t.Fatal()
	}
}

func TestConcurrentTree(t *testing.T) {
	nr := tate.NewNursery(nil)
	var mutex sync.Mutex
	var counter = 0

	nr.Add(func(c *tate.Linker) {

		nrInner1 := tate.NewNursery(c)
		nrInner2 := tate.NewNursery(c)
		nrInner3 := tate.NewNursery(c)

		nrInner1.Add(func(c1 *tate.Linker) {
			for i := 0; i < IterNursery; i++ {
				mutex.Lock()
				counter += 1
				mutex.Unlock()
			}
		})

		nrInner2.Add(func(c2 *tate.Linker) {

			nrInner21 := tate.NewNursery(c2)
			nrInner22 := tate.NewNursery(c2)

			nrInner22.Add(func(c22 *tate.Linker) {
				for i := 0; i < IterNursery; i++ {
					mutex.Lock()
					counter += 1
					mutex.Unlock()
				}
			})

			nrInner21.Add(func(c22 *tate.Linker) {
				for i := 0; i < IterNursery; i++ {
					mutex.Lock()
					counter += 1
					mutex.Unlock()
				}
			})

			for i := 0; i < IterNursery; i++ {
				mutex.Lock()
				counter += 1
				mutex.Unlock()
			}
		})

		nrInner3.Add(func(c3 *tate.Linker) {

			nrInner31 := tate.NewNursery(c3)

			nrInner31.Add(func(c31 *tate.Linker) {
				for i := 0; i < IterNursery; i++ {
					mutex.Lock()
					counter += 1
					mutex.Unlock()
				}
			})

			for i := 0; i < IterNursery; i++ {
				mutex.Lock()
				counter += 1
				mutex.Unlock()
			}
		})
	})
	nr.Join()
	if counter != IterNursery*6 {
		t.Fatal()
	}
}
