package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/smallnest/syncx"
)

type Runnable func()

type FooBar struct {
	n     int
	token *syncx.Token
}

func (fb FooBar) foo(printFoo Runnable) {
	for i := 0; i < fb.n; i++ {
		fb.token.Accquire(context.TODO(), 0)
		printFoo()
		fb.token.Handoff(context.TODO(), 1)
	}
}

func (fb FooBar) bar(printBar Runnable) {
	for i := 0; i < fb.n; i++ {
		fb.token.Accquire(context.TODO(), 1)
		printBar()
		fb.token.Handoff(context.TODO(), 0)
	}
}

func main() {
	printFoo := func() {
		fmt.Print("foo")
	}
	printBar := func() {
		fmt.Print("bar")
	}

	var wg sync.WaitGroup
	wg.Add(2)

	token := syncx.NewToken(2)
	foobar := FooBar{n: 10, token: token}

	go func() {
		foobar.foo(printFoo)
		wg.Done()
	}()
	go func() {
		foobar.bar(printBar)
		wg.Done()
	}()

	// trigger to start
	token.Handoff(context.TODO(), 0)
	wg.Wait()
}
