package main

import (
	"fmt"
)

func main() {
	eo := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for n := range eo {
		if n%2 != 0 {
			fmt.Println(n, ": is odd")
		} else {
			fmt.Println(n, ": is even")
		}
	}
}
