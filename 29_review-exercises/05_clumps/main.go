package main

import "fmt"

func countClumps(xs []int) int {
	clumps := 0
	inClumps := false
	for i := 1; i < len(xs); i++ {
		curr, prev := xs[i], xs[i-1]
		if !inClumps && curr == prev {
			inClumps = true
			clumps++
		}
		if inClumps && curr != prev {
			inClumps = false
		}
	}
	return clumps
}

func main() {
	clumps := countClumps([]int{1, 1, 1, 1, 1})
	fmt.Println(clumps)
}
