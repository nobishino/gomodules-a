package a_test

import a "github.com/nobishino/gomodules-a"

func ExampleA_Print() {
	a := a.A{Number: 10}
	a.Print()
	// Output:
	// This is gomodule-a v1.1: a.Number: 10, c={D:{Number:0} D1_1:{}}
}
