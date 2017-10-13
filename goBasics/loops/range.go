package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

//when iterating over the slice with range 2 values - index, value

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
