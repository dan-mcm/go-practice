package main

import "fmt"

//note const cannot be defined with := notation
const Pi = 3.14

func main() {
	const world = "hassoh"
	fmt.Println("hello", world)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
