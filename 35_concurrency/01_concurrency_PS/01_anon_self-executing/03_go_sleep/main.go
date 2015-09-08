package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		fmt.Println("One")
	}()

	go func() {
		fmt.Println("Two")
	}()

	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}
