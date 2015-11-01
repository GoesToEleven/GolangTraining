package main

import "fmt"

const (
	miTokm  = 1.60934
	pToKg   = 0.453592
	divider = "+------------------------+"
)

func main() {
	fmt.Println("Choose an option")
	fmt.Println("1: Miles to Kilometers")
	fmt.Println("2: Fahrenheit to Celsius")
	fmt.Println("3: Pounds to Kilograms")
	var (
		option int
		number float64
	)
	fmt.Scanf("%d", &option)
	fmt.Print("Enter a number: ")
	fmt.Scanf("%f", &number)

	fmt.Println(divider)
	switch option {
	case 1:
		fmt.Printf("| Miles: %15.2f |\n", number)
		fmt.Println(divider)
		fmt.Printf("| Kilometers: %10.2f |\n", number*miTokm)
	case 2:
		fmt.Printf("| Fahrenheit: %10.2f |\n", number)
		fmt.Println(divider)
		fmt.Printf("| Celsius: %13.2f |\n", (number-32)*5/9)
	case 3:
		fmt.Printf("| Pounds: %14.2f |\n", number)
		fmt.Println(divider)
		fmt.Printf("| Kilograms: %11.2f |\n", number*pToKg)
	}
	fmt.Println(divider)
}
