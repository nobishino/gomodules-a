package a

import (
	"fmt"

	c "github.com/nobishino/gomodules-c/lib"
)

type A struct {
	Number int
}

func (a *A) Print() {
	var c c.C
	fmt.Printf("This is gomodule-a v1.1: a.Number: %d, c=%+v\n", a.Number, c)
}
