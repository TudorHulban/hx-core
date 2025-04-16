package widgets

import (
	hxcomponents "github.com/TudorHulban/hx-core/components"
	hxhtml "github.com/TudorHulban/hx-core/html"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
)

type WidgetCardVerticalInfo struct {
	Title string
	Price string

	ImageSquareSize string
	ImageSource     string

	ActionEndpoint string
}

type ParamsWidgetCardVertical struct {
	WidgetCardVerticalInfo

	CurrencySimbol string
	PriceCaption   string
}

func WidgetCardVertical(params *ParamsWidgetCardVertical) hxprimitives.Node {
	return hxhtml.Div(
		hxprimitives.AttrClass("has-post-thumbnail bw-to-color"),

		hxhtml.Div(
			hxprimitives.AttrClass("post-wrapper"),

			hxhtml.P(
				hxprimitives.AttrClass("post-thumbnail"),

				hxhtml.A(
					hxprimitives.Href(params.ActionEndpoint),

					hxcomponents.Img(
						&hxcomponents.ParamsImage{
							ImageSquareSize: params.ImageSquareSize,
							ImageSource:     params.ImageSource,
							ImageAlt:        params.Title,
						},
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
