package hxcore_test

import (
	"fmt"
	"os"
	"testing"

	hxhtml "github.com/TudorHulban/hx-core/html"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
)

func TestHowToUse(t *testing.T) {
	type person struct {
		Name string
		ID   int
	}

	p1 := person{
		ID:   1,
		Name: "John",
	}

	fragment := hxhtml.Div(
		hxprimitives.AttrIDLength("idsearch-item"),
		hxprimitives.AttrWithValue(
			"style",
			"display: flex; flex-direction: column;",
		),

		hxhtml.H3(
			hxprimitives.Text(
				fmt.Sprintf(
					"Hi %s (%d)!",
					p1.Name,
					p1.ID,
				),
			),
		),
	)

	fmt.Println(
		fragment.Render(os.Stdout),
	)
}
