package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Ints(n)
	fmt.Println(n)
}
