package main

import "fmt"

// Very naive answer.
// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n := 0
	a := 0
	b := 1
	c := a + b
	return func() int {
		var ret int
		switch {
		case n == 0:
			n++
			ret = 0
		case n == 1:
			n++
			ret = 1
		default:
			ret = c
			a = b
			b = c
			c = a + b
		}
		return ret
	}
}

//simplified fibonacci
func fibonacci2() func() int {
    fib := 0
    return func() int {
        first, second := 0, 1
        for i := 0; i < fib; i++ {
            first, second = second, first + second
        }
        fib++
        return first
    }
}

//one based fibonacci
func fibonacci3() func() int {
    last1 := 0
    last2 := 1
    return func() int {
      temp := last1 + last2
      last2, last1 = last1, temp
      return last1
    }
  }

//zero based fibonacci
func fibonacci4() func() int {
    last1 := 0
    last2 := 1
    return func() int {
      last2, last1 = last1, last1 + last2
      return last1 - last2
    }
  }
  
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
