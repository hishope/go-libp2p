package cldr_test

import (
	"fmt"
	"gx/QmVgdgtv5rUcWWcHvHN1vdGkXi8UNztJ7JuT72YUJgG2ic/go-text/cldr"
)

func ExampleSlice() {
	var dr *cldr.CLDR // assume this is initalized

	x, _ := dr.LDML("en")
	cs := x.Collations.Collation
	// remove all but the default
	cldr.MakeSlice(&cs).Filter(func(e cldr.Elem) bool {
		return e.GetCommon().Type != x.Collations.Default()
	})
	for i, c := range cs {
		fmt.Println(i, c.Type)
	}
}