package slices

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		p[i] = make([]uint8, dx)
		for j := 0; j < dy; j++ {
			p[i][j] = uint8((i ^ j))
		}
	}

	return p
}

func Showy() {
	pic.Show(Pic)
}
