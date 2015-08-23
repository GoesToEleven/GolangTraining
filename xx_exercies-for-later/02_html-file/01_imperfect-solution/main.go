package main

import "fmt"

func main() {
	var fname string
	fmt.Print("What is your first name? ")
	fmt.Scan(&fname)
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
