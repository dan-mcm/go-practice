package main

import "fmt"

//use of := short assignment used in place of var declaration with implicity type inside function
//outside function every statement begins with a keyword

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
