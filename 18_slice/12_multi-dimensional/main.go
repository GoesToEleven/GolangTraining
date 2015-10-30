package main

import (
	"fmt"
	"math/rand"
)

func main() {

	transactions := make([][]int, 0, 3)

	for i := 0; i < 3; i++ {
		transaction := make([]int, 0)
		for j := 0; j < int(rand.Float64()*20); j++ {
			transaction = append(transaction, j)
		}
		transactions = append(transactions, transaction)
	}
	fmt.Println(transactions)
}
