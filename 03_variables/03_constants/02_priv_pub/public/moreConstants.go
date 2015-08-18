package public

import "fmt"

const ThisIsPublic = "it is public because it's capitalized"
const thisIsNotPublic = "it is not public because it starts lowercase"

func SomeFunc() {
	fmt.Println(thisIsNotPublic)
}