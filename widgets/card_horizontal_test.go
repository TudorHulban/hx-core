package widgets

import (
	"fmt"
	"testing"

	hxcomponents "github.com/TudorHulban/hx-core/components"
	hxhtml "github.com/TudorHulban/hx-core/html"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
	"github.com/stretchr/testify/require"
)

func TestHorizontalCard(t *testing.T) {
	fragment := WidgetCardHorizontal(
		&ParamsWidgetCardHorizontal{
			WidgetCardHorizontalInfo: WidgetCardHorizontalInfo{
				Title: "Washing Head",
				Price: "40",

				ImageSquareSize: "220",
				ImageSource:     "https://images.pexels.com/photos/668353/pexels-photo-668353.jpeg",

				Text: "Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam",
				Highlights: []*Highlight{
					{
						Label: "Time",
						Text:  "30 min",
					},
					{
						Label: "Equipment",
						Text:  "Shampoo",
					},
				},
			},

			CurrencySimbol: "RON",
			PriceCaption:   "Pret",

			CSSFlexGap: "20px",
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
