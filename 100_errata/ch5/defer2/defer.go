// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 151.

// Defer2 demonstrates a deferred call to runtime.Stack during a panic.
package main

import (
	"fmt"
	"os"
	"runtime"
)

//!+
func main() {
	defer printStack()
	f(3)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

//!-

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

/*
//!+printstack
goroutine 1 [running]:
main.printStack()
	src/gopl.io/ch5/defer2/defer.go:20
main.f(0)
	src/gopl.io/ch5/defer2/defer.go:27
main.f(1)
	src/gopl.io/ch5/defer2/defer.go:29
main.f(2)
	src/gopl.io/ch5/defer2/defer.go:29
main.f(3)
	src/gopl.io/ch5/defer2/defer.go:29
main.main()
	src/gopl.io/ch5/defer2/defer.go:15
//!-printstack
*/
