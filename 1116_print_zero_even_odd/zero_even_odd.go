package main

import (
	"fmt"
	"sync"

	co "github.com/smallnest/leetcode/concurrency"
)

type ZeroEvenOdd struct {
	n   int
	coo *co.Coordinator
}

func (zeo *ZeroEvenOdd) zero(accept co.Accept) {
	for i := 1; i <= zeo.n; i++ {
		zeo.coo.Accquire(0)
		accept(0)
		if i%2 == 0 {
			zeo.coo.Handoff(1)
		} else {
			zeo.coo.Handoff(2)
		}
	}

}

func (zeo *ZeroEvenOdd) even(accept co.Accept) {
	for i := 0; i < zeo.n/2; i++ {
		zeo.coo.Accquire(1)
		accept((i + 1) * 2)
		zeo.coo.Handoff(0)
	}
}

func (zeo *ZeroEvenOdd) odd(accept co.Accept) {
	for i := 0; i < (zeo.n+1)/2; i++ {
		zeo.coo.Accquire(2)
		accept(i*2 + 1)
		zeo.coo.Handoff(0)
	}
}

func main() {
	accept := func(x int) { fmt.Print(x) }

	var wg sync.WaitGroup
	wg.Add(3)

	zeo := &ZeroEvenOdd{
		n:   9,
		coo: co.NewCoordinator(3),
	}
	zeo.coo.Handoff(0)

	go func() {
		zeo.zero(accept)
		wg.Done()
	}()
	go func() {
		zeo.even(accept)
		wg.Done()
	}()
	go func() {
		zeo.odd(accept)
		wg.Done()
	}()

	wg.Wait()
}
