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
	return 0
}

func main() {
	var tr, tco int
	var mod string
	fmt.Scanln(&tr, &tco)
	fmt.Scanln(&mod)
	fmt.Println(tControl(tr, tco, mod))
}
