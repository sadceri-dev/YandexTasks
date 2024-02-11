package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)

	if a == 0 {
		if b == c*c {
			fmt.Println("MANY SOLUTIONS")
		} else {
			fmt.Println("NO SOLUTION")
		}
		return
	}

	x := (c * c) - b
	if x%a != 0 || (c < 0 && b != 0) {
		fmt.Println("NO SOLUTION")
		return
	}

	x /= a
	if x*a+b < 0 || (c == 0 && b < 0) {
		fmt.Println("NO SOLUTION")
	} else {
		fmt.Println(x)
	}
}
