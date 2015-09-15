package main

import (
	"fmt"
	"time"
)

func main() {
	timeAsString := "01/22/2012"
	timeAsTime, _ := time.Parse("01/02/2006", timeAsString)
	fmt.Println(timeAsTime)

	fmt.Println(timeAsTime.Format("01/02/2006"))
	fmt.Println(timeAsTime.Format(time.Kitchen))
	fmt.Println(timeAsTime.Format(time.UnixDate))
}
