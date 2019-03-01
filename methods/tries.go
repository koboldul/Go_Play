package methods

import "fmt"

type MInt int64

func (i *MInt) ToStr() string {

	return "s"
}

func X() {
	var i MInt
	i = 34
	fmt.Printf("%s", i.ToStr())
	fmt.Printf("%v", i)
}
