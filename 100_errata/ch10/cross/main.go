// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 295.

// The cross command prints the values of GOOS and GOARCH for this target.
package main

import (
	"fmt"
	"runtime"
)

//!+
func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}

//!-
