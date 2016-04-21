package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Article struct {
	Name  string
	draft bool
}

func main() {
	myArticle := Article{
		Name:  "Once And Then Again",
		draft: false,
	}

	data, err := json.Marshal(myArticle)
	if err != nil {
		log.Fatalln("couldn't marshall", err.Error())
	} else {
		fmt.Println(string(data))
		os.Stdout.Write(data)
	}
}
