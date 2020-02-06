package main

import (
	"sync"

	co "github.com/smallnest/leetcode/concurrency"
)

type FooBar struct {
	n int
}

func (fb FooBar) foo(printFoo co.Runnable) {
	fooCh <- struct{}{}
	for i := 0; i < fb.n; i++ {
		<-fooCh
		printFoo()
		barCh <- struct{}{}
	}
}

func (fb FooBar) bar(printBar co.Runnable) {
	for i := 0; i < fb.n; i++ {
		<-barCh
		printBar()
		fooCh <- struct{}{}
	}
}

var fooCh = make(chan struct{}, 1)
var barCh = make(chan struct{}, 1)

func main() {
	printFoo := co.WrapPrint("foo")
	printBar := co.WrapPrint("bar")

	var wg sync.WaitGroup
	wg.Add(2)

	foobar := FooBar{n: 10}
	go func() {
		foobar.foo(printFoo)
		wg.Done()
	}()
	go func() {
		foobar.bar(printBar)
		wg.Done()
	}()

	wg.Wait()
}
