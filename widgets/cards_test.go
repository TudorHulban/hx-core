package widgets

import (
	"fmt"
	"testing"

	hxcomponents "github.com/TudorHulban/hx-core/components"
	hxhtml "github.com/TudorHulban/hx-core/html"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
	"github.com/stretchr/testify/require"
)

func TestCards(t *testing.T) {
	img1Src := "https://themes.getmotopress.com/bro-barbershop/wp-content/uploads/sites/64/2024/01/service-1.jpg"
	img2Src := "https://themes.getmotopress.com/bro-barbershop/wp-content/uploads/sites/64/2024/01/service-2.jpg"
	img3Src := "https://themes.getmotopress.com/bro-barbershop/wp-content/uploads/sites/64/2024/01/service-3.jpg"
	img4Src := "https://themes.getmotopress.com/bro-barbershop/wp-content/uploads/sites/64/2024/01/service-4.jpg"

	action1 := "https://themes.getmotopress.com/bro-barbershop/service/washing-head/"
	action2 := "https://themes.getmotopress.com/bro-barbershop/service/beard-trim/"
	action3 := "https://themes.getmotopress.com/bro-barbershop/service/mustache-haircut/"
	action4 := "https://themes.getmotopress.com/bro-barbershop/service/mens-haircut/"

	cards := WidgetCards(
		&ParamsWidgetCards{
			CurrencySimbol: "RON",
			PriceCaption:   "Pret",

			CSSFlexGap: "20px",

			Cards: []*WidgetCardVerticalInfo{
				{
					Title: "Washing Head",
					Price: "49",

					ImageSquareSize: "220",
					ImageSource:     img1Src,

					ActionEndpoint: action1,
				},
				{
					Title: "Beard Trim",
					Price: "39",

					ImageSquareSize: "220",
					ImageSource:     img2Src,

					ActionEndpoint: action2,
				},
				{
					Title: "Mustache Haircut",
					Price: "29",

					ImageSquareSize: "220",
					ImageSource:     img3Src,

					ActionEndpoint: action3,
				},
				{
					Title: "Men's Haircut",
					Price: "59",

					ImageSquareSize: "220",
					ImageSource:     img4Src,

					ActionEndpoint: action4,
				},
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
				hxprimitives.Href("card.css"),
			),
		},

		Body: []hxprimitives.Node{
			cards,
		},
	}

	fmt.Println(
		page.Build().Render(writer),
	)
}
