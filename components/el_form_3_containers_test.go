package hxcomponents

import (
	"fmt"
	"os"
	"testing"

	hxhtml "github.com/TudorHulban/hx-core/html"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
)

func TestNewForm3Containers(t *testing.T) {
	form := NewFormThreeContainers(
		&ParamsNewFormThreeContainers{
			IDDivEnclosing:  "footer-info",
			IDForm:          "form-footer",
			IDDivContainers: "footer-form-info",

			ElementsInputLeft: []hxprimitives.Node{
				hxhtml.Span(
					hxprimitives.Text(
						"Contact and Open Hours",
					),
				),
				hxhtml.Span(
					hxprimitives.Text(
						"+1234567890",
					),
				),
			},

			ElementsInputMiddle: []hxprimitives.Node{
				hxhtml.Span(
					hxprimitives.Text(
						"Address",
					),
				),
				hxhtml.Span(
					hxprimitives.Text(
						"Lorem ipsum 3",
					),
				),
			},

			ElementsInputRight: []hxprimitives.Node{
				hxhtml.Span(
					hxprimitives.Text(
						"Book an appointment",
					),
				),
			},
		},
	)

	fmt.Println(
		form.Render(os.Stdout),
	)
}
