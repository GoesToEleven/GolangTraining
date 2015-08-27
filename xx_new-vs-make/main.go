package main
import (
	"fmt"
	"reflect"
)

func main() {
	scores := make([]int, 10, 100)
	fmt.Println(reflect.TypeOf(scores)) // find the type
	fmt.Println(scores)

	var p *[]int = new([]int)	        // allocates slice structure; *p == nil; rarely useful
	fmt.Printf("%T\n", p) 				// find the type
	fmt.Println(p)


	// Unnecessarily complex:
	var q *[]int = new([]int)
	fmt.Printf("%T\n", q) 				// find the type
	fmt.Println(q)
	*q = make([]int, 100, 100)
	fmt.Printf("%T\n", q) 				// find the type
	fmt.Println(q)


//
//	// Idiomatic:
//	v := make([]int, 100)
}
