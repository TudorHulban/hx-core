package widgets

import (
	"fmt"

	hxcomponents "github.com/TudorHulban/hx-core/components"
	hxhelpers "github.com/TudorHulban/hx-core/helpers"
	hxhtml "github.com/TudorHulban/hx-core/html"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
)

type WidgetCardHorizontalInfo struct {
	Title string
	Price string

	ImageSquareSize string
	ImageSource     string

	Text       string
	Highlights []*Highlight
}

type ParamsWidgetCardHorizontal struct {
	WidgetCardHorizontalInfo

	CurrencySimbol string
	PriceCaption   string

	CSSFlexGap string
}

func WidgetCardHorizontal(params *ParamsWidgetCardHorizontal) hxprimitives.Node {
	return hxhtml.Div(
		hxprimitives.AttrCSS(
			hxhelpers.Sprintf(
				"display:flex;flex-direction:row;gap:%s;",
				params.CSSFlexGap,
			),
		),

		hxprimitives.AttrClass("has-post-thumbnail"),

		hxcomponents.Img(
			&hxcomponents.ParamsImage{
				ImageSquareSize: params.ImageSquareSize,
				ImageSource:     params.ImageSource,
				ImageAlt:        params.Title,
			},
		),

		hxhtml.Div(
			append(
				[]hxprimitives.Node{
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

					hxhtml.Div(
						hxprimitives.AttrClass("content"),

						hxhtml.Span(
							hxprimitives.Text(params.Text),
						),
					),
				},

				hxhelpers.ForEachValueWAddition(
					&hxhelpers.ParamsForEachValueWAddition[*Highlight, hxprimitives.Node]{
						Values: params.Highlights,

						Addition: func() hxprimitives.Node {
							return hxhtml.Div(
								hxprimitives.AttrClass(
									"horizontal-line",
								),
							)
						},

						Process: func(item *Highlight) hxprimitives.Node {
							return hxhtml.Div(
								hxprimitives.Text(
									fmt.Sprintf(
										"%s : %s",
										item.Label,
										item.Text,
									),
								),

								hxhtml.Div(
									hxprimitives.AttrClass(
										"horizontal-line",
									),
								),
							)
						},
					},
				)...,
			)...,
		),
	)
}
