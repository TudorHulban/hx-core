package widgets

import (
	"fmt"
	"testing"

	hxcomponents "github.com/TudorHulban/hx-core/components"
	hxhtml "github.com/TudorHulban/hx-core/html"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
	"github.com/stretchr/testify/require"
)

func TestVerticalCard(t *testing.T) {
	fragment := WidgetCardVertical(
		&ParamsWidgetCardVertical{
			WidgetCardVerticalInfo: WidgetCardVerticalInfo{
				Title: "Jumpin Jack",
				Price: "4000",

				ImageSquareSize: "160",
				ImageSource:     "https://images.pexels.com/photos/19040825/pexels-photo-19040825.jpeg",

				ActionEndpoint: "TBD",
			},

			CurrencySimbol: "RON",
			PriceCaption:   "Price",
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
				hxprimitives.Href("card.css"),
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
