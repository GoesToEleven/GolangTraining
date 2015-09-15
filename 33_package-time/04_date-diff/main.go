package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	from, to := os.Args[1], os.Args[2]

	fromTime, _ := time.Parse("2006-01-02", from)
	toTime, _ := time.Parse("2006-01-02", to)

	dur := toTime.Sub(fromTime)
	fmt.Println("elapsed days:", int(dur/(time.Hour*24)))

}


