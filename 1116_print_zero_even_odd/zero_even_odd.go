package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/smallnest/syncx"
)

type Accept func(x int)

type ZeroEvenOdd struct {
	n     int
	token *syncx.Token
}

func (zeo *ZeroEvenOdd) zero(accept Accept) {
	for i := 1; i <= zeo.n; i++ {
		zeo.token.Accquire(context.TODO(), 0)
		accept(0)
		if i%2 == 0 {
			zeo.token.Handoff(context.TODO(), 1)
		} else {
			zeo.token.Handoff(context.TODO(), 2)
		}
	}

}

func (zeo *ZeroEvenOdd) even(accept Accept) {
	for i := 0; i < zeo.n/2; i++ {
		zeo.token.Accquire(context.TODO(), 1)
		accept((i + 1) * 2)
		zeo.token.Handoff(context.TODO(), 0)
	}
}

func (zeo *ZeroEvenOdd) odd(accept Accept) {
	for i := 0; i < (zeo.n+1)/2; i++ {
		zeo.token.Accquire(context.TODO(), 2)
		accept(i*2 + 1)
		zeo.token.Handoff(context.TODO(), 0)
	}
}

func main() {
	accept := func(x int) { fmt.Print(x) }

	var wg sync.WaitGroup
	wg.Add(3)

	token := syncx.NewToken(3)
	zeo := &ZeroEvenOdd{
		n:     9,
		token: token,
	}

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

	// trigger to start
	token.Handoff(context.TODO(), 0)
	wg.Wait()
}
