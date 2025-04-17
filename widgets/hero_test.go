package widgets

import (
	"fmt"
	"testing"

	hxcomponents "github.com/TudorHulban/hx-core/components"
	hxhtml "github.com/TudorHulban/hx-core/html"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
	"github.com/stretchr/testify/require"
)

func TestHero(t *testing.T) {
	fragment := WidgetHero(
		&ParamsWidgetHero{
			Title:   "Expert cuts and classic styles",
			Message: "experience the finest grooming services in town. from traditional haircuts to modern styling, we've got you covered.",

			ButtonPrimaryInfo: hxcomponents.ParamsElementARef{
				CSSClass: "button primary",
				Caption:  "Book appointment",
			},
			ButtonSecondaryInfo: hxcomponents.ParamsElementARef{
				CSSClass: "button secondary",
				Caption:  "View services",
			},

			ParamsImage: hxcomponents.ParamsImage{
				ImageSquareSize: "400",
				ImageAlt:        "Barber Shop",
				ImageSource:     "https://themes.getmotopress.com/bro-barbershop/wp-content/uploads/sites/64/2024/01/appointment-img.jpg",
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
				hxprimitives.Href("hero.css"),
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
