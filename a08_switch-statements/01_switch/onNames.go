package main

import "fmt"

/*
  no default fallthrough
  fallthrough is optional
  -- you can specify fallthrough by explicitly stating it
  -- break isn't needed like in other languages
 */

func main() {
	switch "Medhi" {
	case "Daniel":
		fmt.Println("Wassup Jenny")
	case "Medhi":
		fmt.Println("Wassup Medhi")
	case "Jenny":
		fmt.Println("Wassup Sushant")
	default:
		fmt.Println("Have you no friends?")
	}
}
