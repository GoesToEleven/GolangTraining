package main

import (
	"fmt"
	"os"
)

func main() {

	name := os.Args[1]

	fmt.Println(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Go Rocks!</title>
	</head>
	<body>
	<h1>` +
		name +
		`</h1>
	</body>
	</html>
	`)
}

/*
go build
./02_string-to-html Todd > index.html
*/
