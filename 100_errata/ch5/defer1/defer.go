// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 150.

// Defer1 demonstrates a deferred call being invoked during a panic.
package main

import "fmt"

//!+f
func main() {
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

//!-f

/*
//!+stdout
f(3)
f(2)
f(1)
defer 1
defer 2
defer 3
//!-stdout

//!+stderr
panic: runtime error: integer divide by zero
main.f(0)
        src/gopl.io/ch5/defer1/defer.go:14
main.f(1)
        src/gopl.io/ch5/defer1/defer.go:16
main.f(2)
        src/gopl.io/ch5/defer1/defer.go:16

main.f(3)
        src/gopl.io/ch5/defer1/defer.go:16
main.main()
        src/gopl.io/ch5/defer1/defer.go:10
//!-stderr
*/
