package methods

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)

	if n > 0 {
		for i, _ := range b[:n] {
			switch {
			case b[i] > 'm' && b[i] <= 'z':
				b[i] -= 13
			case b[i] <= 'Z' && b[i] > 'M':
				b[i] -= 13
			default:
				b[i] += 13
			}
		}
	}
	return n, err
}

func Xr() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
