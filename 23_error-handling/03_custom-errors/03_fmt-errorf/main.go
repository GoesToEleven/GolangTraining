package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := Sqrt(-10.23)
	if err != nil {
		log.Fatalln(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("norgate math again: square root of negative number: %v", f)
	}
	// implementation
	return 42, nil
}
