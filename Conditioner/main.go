package main

import "fmt"

func tControl(troom, tcond int, mode string) int {
	if mode == "heat" {
		if troom < tcond {
			return tcond
		} else {
			return troom
		}
	} else if mode == "freeze" {
		if troom > tcond {
			return tcond
		} else {
			return troom
		}
	} else if mode == "auto" {
		return tcond
	} else if mode == "fan" {
		return troom
	}

	// Default return statement
	return 0
}

func main() {
	fmt.Println(tControl(30, 20, "fan")) // 20
}
