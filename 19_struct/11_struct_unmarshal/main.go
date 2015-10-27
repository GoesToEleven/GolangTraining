package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Email    string
	Loggedin string
	Pictures []string
}

func main() {
	var m Data
	fmt.Println(m)
	fmt.Println(m.Email)
	fmt.Println(m.Loggedin)
	fmt.Println(m.Pictures)

	str := `{"Email":"tm@gmail.com","Loggedin":"true","Pictures":["Crimson.jpg","Red.jpg","Ruby.jpg","Maroon.jpg"]}`
	bs := []byte(str)
	json.Unmarshal(bs, &m)

	fmt.Println(m)
	fmt.Println(m.Email)
	fmt.Println(m.Loggedin)
	fmt.Println(m.Pictures)
}
