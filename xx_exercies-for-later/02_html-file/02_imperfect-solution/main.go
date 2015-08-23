package main

import (
	"fmt"
	"os"
)

func main() {
	fname := os.Args[1]
	fmt.Println(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title></title>
	</head>
	<body>
	`)
	fmt.Print(`<h1>`)
	fmt.Print(fname)
	fmt.Println(`</h1>`)
	fmt.Println(`
	</body>
	</html>
	`)
}
