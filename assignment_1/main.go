package main

import "fmt"

func main() {
	var ints []int
	for i := 0; i <= 10; i++ {
		ints = append(ints, i)
	}
	for _, i := range ints {
		if i%2 == 0 {
			fmt.Printf("%d is even\n", i)
		} else {
			fmt.Printf("%d is odd\n", i)
		}
	}
}
