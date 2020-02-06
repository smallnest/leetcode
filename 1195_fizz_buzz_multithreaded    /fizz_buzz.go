package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	co "github.com/smallnest/leetcode/concurrency"
)

type FizzBuzz struct {
	n   int
	coo *co.Coordinator
}

func (fb *FizzBuzz) fizz(printFizz co.Accept) {
	for i := 1; i <= fb.n; i++ {
		fb.coo.Accquire(0)
		if i%3 == 0 && i%15 != 0 {
			printFizz(i)
		}
		fb.coo.Handoff(1)
	}
}

func (fb *FizzBuzz) buzz(printBuzz co.Accept) {
	for i := 1; i <= fb.n; i++ {
		fb.coo.Accquire(1)
		if i%5 == 0 && i%15 != 0 {
			printBuzz(i)
		}
		fb.coo.Handoff(2)
	}
}

func (fb *FizzBuzz) fizzbuzz(printFizzBuzz co.Accept) {
	for i := 1; i <= fb.n; i++ {
		fb.coo.Accquire(2)
		if i%15 == 0 {
			printFizzBuzz(i)
		}
		fb.coo.Handoff(3)
	}
}

func (fb *FizzBuzz) number(printNumber co.Accept) {
	for i := 1; i <= fb.n; i++ {
		fb.coo.Accquire(3)
		if i%3 != 0 && i%5 != 0 {
			printNumber(i)
		}
		fb.coo.Handoff(0)
	}
}

func main() {
	if len(os.Args) == 1 || os.Args[1] == "1" {
		main1()
		return
	}
}
func main1() {
	var result []string
	fizz := func(x int) { result = append(result, "fizz") }
	buzz := func(x int) { result = append(result, "buzz") }
	fizzbuzz := func(x int) { result = append(result, "fizzbuzz") }
	number := func(x int) { result = append(result, strconv.Itoa(x)) }

	var wg sync.WaitGroup
	wg.Add(4)

	fb := FizzBuzz{
		n:   15,
		coo: co.NewCoordinator(4),
	}

	fb.coo.Handoff(0)

	go func() {
		fb.fizz(fizz)
		wg.Done()
	}()

	go func() {
		fb.buzz(buzz)
		wg.Done()
	}()

	go func() {
		fb.fizzbuzz(fizzbuzz)
		wg.Done()
	}()

	go func() {
		fb.number(number)
		wg.Done()
	}()

	wg.Wait()

	fmt.Println(strings.Join(result, ", "))
}
