package hxcomponents

import (
	"fmt"
	"testing"
)

func TestARef(t *testing.T) {
	element := ElementARef(
		&ParamsElementARef{
			Route:   "/home",
			Caption: "Home",
		},
	)

	fmt.Println(
		element,
	)
}
