package math

import (
	"fmt"
	"testing"
	"testing/quick"
)

func TestAdder(t *testing.T) {
	type args struct {
		xs []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Simple test case. Adding 4 and 7.",
			args: args{xs: []int{4, 7}},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Adder(tt.args.xs...); got != tt.want {
				t.Errorf("Adder() = %v, want %v", got, tt.want)
			}
		})
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
