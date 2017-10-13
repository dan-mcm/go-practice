package main

func main() {
	var s []int
	printSlice(s)

	//append works on nil slices
	s = append(s, 0)
	printSlice(s)

	//the slice grows as needed
	s = append(s, 1)
	printSlice(s)

	//we can add more than one elemnet at a time
	s = append(s, 2, 3, 4)
	printSlice(s)
}
