package hxcore_test

import (
	"fmt"
	"os"
	"testing"

	hxcomponents "github.com/TudorHulban/hx-core/components"
	hxhtml "github.com/TudorHulban/hx-core/html"
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
		hxcomponents.AttrIDLength("idsearch-item"),
		hxcomponents.AttrWithValue(
			"style",
			"display: flex; flex-direction: column;",
		),

		hxhtml.H3(
			hxcomponents.Text(
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
