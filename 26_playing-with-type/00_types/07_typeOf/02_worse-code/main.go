package main

import (
	"fmt"
	"reflect"
)

var a string = "this is stored in the variable a" // package scope
var b, c string = "stored in b", "stored in c"    // package scope
var d string                                      // package scope

func main() {

	d = "stored in d" // declaration above; assignment here; package scope
	var e int = 42    // function scope - subsequent variables have same package scope:
	f := 43
	g := "stored in g"
	h, i := "stored in h", "stored in i"
	j, k, l, m := 44.7, true, false, 'm' // single quotes
	n := "n"                             // double quotes
	o := `o`                             // back ticks

	fmt.Println("a - ", reflect.TypeOf(a), " - ", a)
	fmt.Println("b - ", reflect.TypeOf(b), " - ", b)
	fmt.Println("c - ", reflect.TypeOf(c), " - ", c)
	fmt.Println("d - ", reflect.TypeOf(d), " - ", d)
	fmt.Println("e - ", reflect.TypeOf(e), " - ", e)
	fmt.Println("f - ", reflect.TypeOf(f), " - ", f)
	fmt.Println("g - ", reflect.TypeOf(g), " - ", g)
	fmt.Println("h - ", reflect.TypeOf(h), " - ", h)
	fmt.Println("i - ", reflect.TypeOf(i), " - ", i)
	fmt.Println("j - ", reflect.TypeOf(j), " - ", j)
	fmt.Println("k - ", reflect.TypeOf(k), " - ", k)
	fmt.Println("l - ", reflect.TypeOf(l), " - ", l)
	fmt.Println("m - ", reflect.TypeOf(m), " - ", m)
	fmt.Println("n - ", reflect.TypeOf(n), " - ", n)
	fmt.Println("o - ", reflect.TypeOf(o), " - ", o)
	fmt.Printf("o -  %T\n", o)
}
