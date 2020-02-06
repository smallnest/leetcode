package co

import "fmt"

// Runnable is a simple function.
type Runnable func()

// Accept accepts a integer and print something.
type Accept func(x int)

// WrapPrint wraps `fmt.Print(s)` as Runnable.
func WrapPrint(s string) Runnable {
	return func() {
		fmt.Print(s)
	}
}

// WrapPrintln wraps `fmt.Println(s)` as Runnable.
func WrapPrintln(s string) Runnable {
	return func() {
		fmt.Println(s)
	}
}
