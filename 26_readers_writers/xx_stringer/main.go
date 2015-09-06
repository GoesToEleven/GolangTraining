package main
import "fmt"

type person struct {
	name string
	surname string
}

func (p person) String() string {
	return p.name + " " + p.surname
}

func main() {
	p1 := person{"Todd", "McLeod"}
	fmt.Println(p1.String())
}
