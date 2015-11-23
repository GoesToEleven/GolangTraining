package main
import (
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

func main() {
	p := "mywifesnameandbirthday"
	bs, _ := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	fmt.Println(bs)
	fmt.Println(string(bs))
	fmt.Printf("%x",bs)

	err := bcrypt.CompareHashAndPassword(bs, []byte(p))
	if err != nil {
		fmt.Println("Doesn't match")
	} else {
		fmt.Println("match")
	}
}
