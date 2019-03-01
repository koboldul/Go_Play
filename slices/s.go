package slices

import "fmt"

func Testy() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
	fmt.Println(primes)
}

func Bitls() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[2:4]
	fmt.Println(a, b)

	b[0] = "XXX"
	a = names[0:3]
	fmt.Println(a, b)
	fmt.Println(names)
}

func AnomArray() {
	d := []struct {
		X int
		B bool
	}{
		{1, true},
		{3, false},
		{2, true},
	}

	for i, v := range d {
		fmt.Printf("%v at %d \n", v, i)
		d[i].X = 555 + i
	}

	for i, v := range d {
		fmt.Printf("%v at %d \n", v, i)
	}
}

func Reslice() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
