package main

import (
	"fmt"
)

func triangle(a, b, c int, res string) string {
	if a+b > c && a+c > b && b+c > a {
		res = "YES"
	} else {
		res = "NO"
	}
	return res
}
func main() {
	var a, b, c int
	var res string
	fmt.Scan(&a, &b, &c)
	fmt.Println(triangle(a, b, c, res))
}
