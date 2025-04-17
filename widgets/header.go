package widgets

import (
	hxcomponents "github.com/TudorHulban/hx-core/components"
	hxhtml "github.com/TudorHulban/hx-core/html"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
)

type ParamsWidgetHeader struct {
	Title string

	hxcomponents.ParamsImage
}

func WidgetHeader(params *ParamsWidgetHeader) hxprimitives.Node {
	return hxhtml.Div(
		hxprimitives.AttrClass("page-header-info"),

		hxhtml.H1(
			hxprimitives.AttrClass("entry-title"),
			hxprimitives.Text(params.Title),
		),

		hxhtml.Div(
			hxprimitives.AttrClass("header-images"),

			hxhtml.Img(
				hxprimitives.AttrWithValue("src", params.ImageSource),
				hxprimitives.AttrWithValue("alt", params.ImageAlt),
			),
		),
	)
}
