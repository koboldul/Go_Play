package methods

import "fmt"

type Abser interface {
	Abs() float64
}

type Verty struct {
	C   int
	ani interface{}
}

func (v Verty) Abs() float64 {
	return float64(v.C)
}

func Xi() {
	var a Abser
	v := Verty{333, "ss"}
	fmt.Printf("%v --> %T \n", v, v)

	v.ani = 567
	a = v
	fmt.Println(a)
}
