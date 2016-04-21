package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	kmToMi = 0.621371
	ftToMi = 0.000189394
	mToMi  = 0.000621371
	miToKm = 1.60934
	miToFt = 5280
	miToM  = 1609.34
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("Not enough arguments")
	}

	from := os.Args[1]
	to := os.Args[2]

	var (
		fromNum  float64
		fromType string
		err      error
		miles    float64
		result   float64
	)

	switch {
	case strings.HasSuffix(from, "mi"):
		fromNum, err = strconv.ParseFloat(from[:len(from)-2], 64)
		miles = fromNum
		fromType = "mi"
	case strings.HasSuffix(from, "km"):
		fromNum, err = strconv.ParseFloat(from[:len(from)-2], 64)
		miles = fromNum * kmToMi
		fromType = "km"
	case strings.HasSuffix(from, "ft"):
		fromNum, err = strconv.ParseFloat(from[:len(from)-2], 64)
		miles = fromNum * ftToMi
		fromType = "ft"
	case strings.HasSuffix(from, "m"):
		fromNum, err = strconv.ParseFloat(from[:len(from)-1], 64)
		miles = fromNum * mToMi
		fromType = "m"
	default:
		log.Fatalf("Unidentified type detected on: %s\n", from)
	}

	if err != nil {
		log.Fatalln(err)
	}

	switch to {
	case "mi":
		result = miles
	case "km":
		result = miles * miToKm
	case "ft":
		result = miles * miToFt
	case "m":
		result = miles * miToM
	default:
		log.Fatalf("Unidentified type detected on: %s\n", to)
	}

	fmt.Println("<!DOCTYPE html>")
	fmt.Println("<html>")
	fmt.Println("\t<head></head>")
	fmt.Println("\t<body>")
	switch fromType {
	case "mi":
		fmt.Printf("\t\tMiles: %.2f<br>\n", fromNum)
	case "km":
		fmt.Printf("\t\tKilometers: %.2f<br>\n", fromNum)
	case "ft":
		fmt.Printf("\t\tFeet: %.2f<br>\n", fromNum)
	case "m":
		fmt.Printf("\t\tMeters: %.2f<br>\n", fromNum)
	}
	switch to {
	case "mi":
		fmt.Printf("\t\tMiles: %.2f<br>\n", result)
	case "km":
		fmt.Printf("\t\tKilometers: %.2f<br>\n", result)
	case "ft":
		fmt.Printf("\t\tFeet: %.2f<br>\n", result)
	case "m":
		fmt.Printf("\t\tMeters: %.2f<br>\n", result)
	}
	fmt.Println("\t</body>")
	fmt.Println("</html>")
}
