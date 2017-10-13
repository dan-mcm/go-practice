package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	//deletes map value
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	//if key is in m, ok is true
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
