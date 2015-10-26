package main
import "fmt"

func main() {
	var x [256]byte
	fmt.Println(len(x))
	fmt.Println(x[0])
	var i int64
	for i = 0; i < 256; i++ {
		x[i] = byte(i)
	}
	for _, v := range x {
		fmt.Printf("%v - %T - %b\n", v, v, v)
	}
}
