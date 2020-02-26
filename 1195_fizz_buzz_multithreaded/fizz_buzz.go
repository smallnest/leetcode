package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/smallnest/syncx"
)

type Accept func(x int)

type FizzBuzz struct {
	n     int
	token *syncx.Token
}

func (fb *FizzBuzz) fizz(printFizz Accept) {
	for i := 1; i <= fb.n; i++ {
		fb.token.Accquire(context.TODO(), 0)
		if i%3 == 0 && i%15 != 0 {
			printFizz(i)
		}
		fb.token.Handoff(context.TODO(), 1)
	}
}

func (fb *FizzBuzz) buzz(printBuzz Accept) {
	for i := 1; i <= fb.n; i++ {
		fb.token.Accquire(context.TODO(), 1)
		if i%5 == 0 && i%15 != 0 {
			printBuzz(i)
		}
		fb.token.Handoff(context.TODO(), 2)
	}
}

func (fb *FizzBuzz) fizzbuzz(printFizzBuzz Accept) {
	for i := 1; i <= fb.n; i++ {
		fb.token.Accquire(context.TODO(), 2)
		if i%15 == 0 {
			printFizzBuzz(i)
		}
		fb.token.Handoff(context.TODO(), 3)
	}
}

func (fb *FizzBuzz) number(printNumber Accept) {
	for i := 1; i <= fb.n; i++ {
		fb.token.Accquire(context.TODO(), 3)
		if i%3 != 0 && i%5 != 0 {
			printNumber(i)
		}
		fb.token.Handoff(context.TODO(), 0)
	}
}

func main() {
	var result []string
	fizz := func(x int) { result = append(result, "fizz") }
	buzz := func(x int) { result = append(result, "buzz") }
	fizzbuzz := func(x int) { result = append(result, "fizzbuzz") }
	number := func(x int) { result = append(result, strconv.Itoa(x)) }

	var wg sync.WaitGroup
	wg.Add(4)

	token := syncx.NewToken(4)
	fb := FizzBuzz{
		n:     15,
		token: token,
	}

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

	// trigger to start
	token.Handoff(context.TODO(), 0)
	wg.Wait()

	fmt.Println(strings.Join(result, ", "))
}
