package main

import (
	"fmt"
	"time"
)

func main() {
	timeAsString := "2012-01-22"
	fmt.Println(time.Parse("2006-01-01_this-does-not-compile", timeAsString))
}
