package main

import "fmt"

//initializing multiple variables of same type
var i, j int = 1, 2

func main() {
	//initializing multiple variables of different types
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
