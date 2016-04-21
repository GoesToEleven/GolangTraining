package main

import(
	"fmt"
	"sync"
	"time"
)



func main() {
	var count int64
	c := make(chan int64)
	var wg sync.WaitGroup

	// bees
	for i:=0; i<5000;i++{
		wg.Add(1)
		go func(in chan int64) {
			defer wg.Done()
			time.Sleep(100)
			in <- 2
		}(c)
	}

	// hive
	go func() {
		for out := range c {
			count += out
		}
	}()

	wg.Wait()
	// bang! but why?
	fmt.Println(count)
}
