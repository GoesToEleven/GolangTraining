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
		ErrNorgateMath := fmt.Errorf("norgate math again: square root of negative number: %v", f)
		return 0, ErrNorgateMath
	}
	// implementation
	return 42, nil
}
