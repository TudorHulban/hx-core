package widgets

import (
	"fmt"
	"testing"
	"time"

	hxcomponents "github.com/TudorHulban/hx-core/components"
	hxhtml "github.com/TudorHulban/hx-core/html"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
	"github.com/stretchr/testify/require"
)

func TestWidgetInputDate(t *testing.T) {
	fragment := WidgetInputDate(
		&ParamsWidgetInputDate{
			CSSID: "schedule",

			DateValue:   time.Now(),
			HowManyDays: 3,
		},
	)

	writer, errWriter := getFileWriter(t.Name() + ".html")
	require.NoError(t, errWriter)

	defer writer.Close()

	page := hxcomponents.Page{
		Title: t.Name(),

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
				hxprimitives.Href("https://cdn.jsdelivr.net/npm/flatpickr/dist/flatpickr.min.css"),
			),
			hxhtml.Link(
				hxprimitives.Rel("stylesheet"),
				hxprimitives.Href("https://npmcdn.com/flatpickr/dist/themes/dark.css"),
			),
			hxhtml.Link(
				hxprimitives.Rel("stylesheet"),
				hxprimitives.Href("input_date.css"),
			),
		},

		Body: []hxprimitives.Node{
			fragment.LinkJavascript,
			fragment.HTML,
		},
	}

	fmt.Println(
		page.Build().Render(writer),
	)
}
