package main

import (
	"fmt"
	"github.com/GoesToEleven/GolangTraining/02_package/stringutil"
	"github.com/GoesToEleven/GolangTraining/02_package/icomefromalaska"
	//someAlias "github.com/GoesToEleven/GolangTraining/02_package/icomefromalaska"
	/*for local imports 
	this works only after using ``` go mod init 02_package ``` in the 02_package directory
	//02_package/stringutil
	//02_package/icomefromalaska
	// the local imports works as fine as the github imports and are even shorter and easier to understand
	
)

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Println(stringutil.MyName)
	fmt.Println(winniepooh.BearName)
}
