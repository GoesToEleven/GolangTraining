package main

import (
	"fmt"

	"github.com/ttacon/chalk"
)

func main() {
	strange := chalk.Red.NewStyle()
	strange.WithBackground(chalk.Blue).WithTextStyle(chalk.Underline)
	fmt.Println(strange.Style("Hello Color World!"))
}
