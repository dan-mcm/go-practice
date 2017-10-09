package main

import "fmt"

//without explicit type variable's type is inferred
func main() {
	v := 42
	fmt.Printf("v is of type %T\n", v)
}
