package main

import "fmt"

//note variable type comes after declaration of variable name
func add(x int, y int) int {
	return x + y
}

//note you can declare multiple variables of same types like so:
func modifiedAdd(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(add(42, 13))
}
