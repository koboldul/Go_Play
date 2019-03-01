package closures

import "fmt"

type Point struct {
	X int
	Y int
}

func CheckRef() func() {
	p := Point{1, 2}

	var i = 34

	f := func() {
		fmt.Println(p)
		fmt.Println(&i)
	}

	p.X = 23

	fmt.Println(&i)

	return f
}
