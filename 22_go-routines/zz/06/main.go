package main

import (
	"fmt"
)

func main() {
	xs := []string{"Foo:", "Bar:"}
	xc := incrementor(xs)

	var sum int
	for n := range puller(xc) {
		sum += n
	}
	fmt.Println("Final Counter:", sum)
}

func incrementor(xs []string) []<-chan int {
	var xc []<-chan int
	for _, s := range xs {
		out := make(chan int)
		go func(id string){
			for i := 0; i < 20; i++ {
				out <- 1
				fmt.Println(id, i)
			}
			close(out)
		}(s)
		xc = append(xc, out)
	}
	return xc
}

func puller(xc []<-chan int) chan int {
	out := make(chan int)
	for _, c := range xc {
		go func(ch <-chan int){
			var sum int
			for n := range ch {
				sum += n
			}
			out <- sum
		}(c)
	}
	close(out)
	return out
}

// go run -race main.go
// vs
// go run main.go
