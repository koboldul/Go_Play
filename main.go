package main

import (
	"fmt"

	"./closures"
	"./machine_m"
	"./methods"
	"./slices"
)

func main() {
	var choose string
	choose = "mt"

	switch {
	case choose == "c":
		f := closures.CheckRef()
		f()
	case choose == "mt":
		methods.Xr()
	case choose == "f":
		closures.X()
	case choose == "s":
		slices.Reslice()
	case choose == "j":
		slices.Showy()
	case choose == "m":
		machine_m.Testy()
	default:
		fmt.Println("Nothing")
	}

}
