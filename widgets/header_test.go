package widgets

import (
	"fmt"
	"testing"

	hxcomponents "github.com/TudorHulban/hx-core/components"
	hxhtml "github.com/TudorHulban/hx-core/html"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
	"github.com/stretchr/testify/require"
)

func TestHeader(t *testing.T) {
	fragment := WidgetHeader(
		&ParamsWidgetHeader{
			Title:             "Washing Head",
			BreadcrumbCaption: "Barbershop/Washing Head",

			ParamsImage: hxcomponents.ParamsImage{
				ImageSquareSize: "160",
				ImageAlt:        "Washing Head",
				ImageSource:     "https://themes.getmotopress.com/bro-barbershop/wp-content/uploads/sites/64/2024/01/hero-bg-1.jpg",
			},
		},
	)

	writer, errWriter := getFileWriter(t.Name() + ".html")
	require.NoError(t, errWriter)

	defer writer.Close()

	page := hxcomponents.Page{
		Head: []hxprimitives.Node{
			hxhtml.Link(
				hxprimitives.Rel("stylesheet"),
				hxprimitives.Href("css_base.css"),
			),
			hxhtml.Link(
				hxprimitives.Rel("stylesheet"),
				hxprimitives.Href("css_site.css"),
			),
			hxhtml.Link(
				hxprimitives.Rel("stylesheet"),
				hxprimitives.Href("header.css"),
			),
		},

		Body: []hxprimitives.Node{
			fragment,
		},
	}

	fmt.Println(
		page.Build().Render(writer),
	)
}
