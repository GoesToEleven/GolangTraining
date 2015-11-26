package main

import "fmt"

var a = "this is stored in the variable a" // package scope
var b, c string = "stored in b", "stored in c"    // package scope
var d string                                      // package scope

func main() {

	d = "stored in d" // declaration above; assignment here; package scope
	var e = 42    // function scope - subsequent variables have same package scope:
	f := 43
	g := "stored in g"
	h, i := "stored in h", "stored in i"
	j, k, l, m := 44.7, true, false, 'm' // single quotes
	n := "n"                             // double quotes
	o := `o`                             // back ticks

	fmt.Println("a - ", a)
	fmt.Println("b - ", b)
	fmt.Println("c - ", c)
	fmt.Println("d - ", d)
	fmt.Println("e - ", e)
	fmt.Println("f - ", f)
	fmt.Println("g - ", g)
	fmt.Println("h - ", h)
	fmt.Println("i - ", i)
	fmt.Println("j - ", j)
	fmt.Println("k - ", k)
	fmt.Println("l - ", l)
	fmt.Println("m - ", m)
	fmt.Println("n - ", n)
	fmt.Println("o - ", o)
}
