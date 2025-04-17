package widgets

import (
	hxhelpers "github.com/TudorHulban/hx-core/helpers"
	hxhtml "github.com/TudorHulban/hx-core/html"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
)

type ParamsWidgetCards struct {
	Title string

	CurrencySimbol string
	PriceCaption   string

	CSSFlexGap string

	Cards []*WidgetCardVerticalInfo
}

func WidgetCards(params *ParamsWidgetCards) hxprimitives.Node {
	return hxhtml.Div(
		hxhtml.Div(
			hxprimitives.AttrClass(
				"services-list-title",
			),

			hxhtml.H2(
				hxprimitives.AttrClass(
					"item-title",
				),

				hxprimitives.Text(
					params.Title,
				),
			),
		),

		hxhtml.Div(
			hxprimitives.AttrClass(
				"centered-services-list",
			),

			hxhtml.Div(
				append(
					[]hxprimitives.Node{
						hxprimitives.AttrCSS(
							hxhelpers.Sprintf(
								"display:flex;flex-direction:row;gap:%s;",

								params.CSSFlexGap,
							),
						),

						hxprimitives.AttrClass(
							"services-list",
						),
					},

					hxhelpers.ForEachValue(
						params.Cards,

						func(item *WidgetCardVerticalInfo) hxprimitives.Node {
							return WidgetCardVertical(
								&ParamsWidgetCardVertical{
									WidgetCardVerticalInfo: *item,

									CurrencySimbol: params.CurrencySimbol,
									PriceCaption:   params.PriceCaption,
								},
							)
						},
					)...,
				)...,
			),
		),
	)
}
