package math

import (
	"fmt"
	"testing"
	"testing/quick"
)

func TestAdder(t *testing.T) {
	result := Adder(4, 7)
	if result != 11 {
		t.Fatal("4 + 7 did not equal 11")
	}
}

func BenchmarkAdder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Adder(4, 7)
	}
}

func ExampleAdder() {
	fmt.Println(Adder(4, 7))
	// Output:
	// 11
}

func ExampleAdder_multiple() {
	fmt.Println(Adder(3, 6, 7, 4, 61))
	// Output:
	// 81
}

func TestAdderBlackbox(t *testing.T) {
	err := quick.Check(a, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func a(x, y int) bool {
	return Adder(x, y) == x+y
}
