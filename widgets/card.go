package widgets

import (
	hxhtml "github.com/TudorHulban/hx-core/html"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
)

type WidgetCardInfo struct {
	Title string
	Price string

	ImageSquareSize string
	ImageSource     string

	ActionEndpoint string
}

type ParamsWidgetCard struct {
	WidgetCardInfo

	CurrencySimbol string
	PriceCaption   string
}

func WidgetCard(params *ParamsWidgetCard) hxprimitives.Node {
	return hxhtml.Div(
		hxprimitives.AttrClass("has-post-thumbnail bw-to-color"),

		hxhtml.Div(
			hxprimitives.AttrClass("post-wrapper"),

			hxhtml.P(
				hxprimitives.AttrClass("post-thumbnail"),

				hxhtml.A(
					hxprimitives.Href(params.ActionEndpoint),

					hxhtml.Img(
						hxprimitives.AttrWithValue("sizes", "auto, (max-width: 920px) 100vw, 920px"),
						hxprimitives.AttrWithValue("decoding", "async"),
						hxprimitives.AttrWithValue("height", params.ImageSquareSize),
						hxprimitives.AttrWithValue("width", params.ImageSquareSize),
						hxprimitives.AttrWithValue("alt", params.Title),
						hxprimitives.AttrWithValue("loading", "lazy"),
						hxprimitives.AttrWithValue("src", params.ImageSource),
						hxprimitives.AttrClass("attachment-post-thumbnail"),
					),
				),
			),

			hxhtml.Div(
				hxprimitives.AttrClass("post-title-wrapper"),

				hxhtml.H2(
					hxprimitives.AttrClass("entry-title post-title"),

					hxhtml.A(
						hxprimitives.Href(params.ActionEndpoint),
						hxprimitives.Text(params.Title),
					),
				),

				hxhtml.P(
					hxprimitives.AttrClass("service-price"),

					hxhtml.Span(
						hxprimitives.AttrClass("price-title"),
						hxprimitives.Text(params.PriceCaption+":"),
					),

					hxhtml.Span(
						hxprimitives.AttrClass("price"),

						hxhtml.Span(
							hxprimitives.AttrClass("currency"),
							hxprimitives.Text(params.CurrencySimbol),
						),

						hxprimitives.Text(params.Price),
					),
				),
			),
		),
	)
}
