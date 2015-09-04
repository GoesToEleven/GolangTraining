package main
import "fmt"

func main() {
	var y int = 7
	yPtr := new(int)
	*yPtr = y
	fmt.Println(yPtr)
	fmt.Println(*yPtr)
}
