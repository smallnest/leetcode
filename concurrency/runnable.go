package co

import "fmt"

// Runnable is a simple function.
type Runnable func()

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
