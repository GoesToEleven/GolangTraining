package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

type abk_int int
type hotdog int

var x_myint abk_int

func main() {

//	exercise3()
//	exercise4()
	exercise5()

}

func exercise3() {

	s := fmt.Sprintf("%d %s %t",x,y,z)
	fmt.Println(s)
}

func exercise4() {
	fmt.Printf("%d %T\n",x_myint,x_myint)
	x_myint = abk_int(100)

	fmt.Printf("%d %T\n",x_myint,x_myint)
}

func exercise5() {
	fmt.Printf("%d %T\n",x_myint,x_myint)
	x_myint = abk_int(100)
	fmt.Printf("%d %T\n",x_myint,x_myint)

	x = int(x_myint)
	fmt.Printf("%d %T\n",x, x)

}
