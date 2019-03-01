package closures

import "fmt"

func fibonacci() func() int {
	i := 0
	j := 1

	return func() int {
		t := i
		i = j
		j = t + j

		return j
	}
}

func X() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
