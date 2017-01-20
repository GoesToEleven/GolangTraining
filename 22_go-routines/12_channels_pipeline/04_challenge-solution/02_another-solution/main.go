package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const numFactorials = 100
const rdLimit = 20

func main() {
	var w sync.WaitGroup
	w.Add(numFactorials)
	factorial(numFactorials, &w)
	w.Wait()

}

func factorial(n int, wmain *sync.WaitGroup) {
	var w sync.WaitGroup
	rand.Seed(42)

	w.Add(n + 1)

	for j := 1; j <= n; j++ {

		go func() {
			f := rand.Intn(rdLimit)
			w.Wait()
			total := 1
			for i := f; i > 0; i-- {

				total *= i
			}
			fmt.Println(f, total)
			(*wmain).Done()
			//out <- total

		}()
		w.Done()
	}
	fmt.Println("All done with initialization")
	w.Done()
}