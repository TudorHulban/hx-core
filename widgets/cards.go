package widgets

import (
	hxhelpers "github.com/TudorHulban/hx-core/helpers"
	hxhtml "github.com/TudorHulban/hx-core/html"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
)

type ParamsWidgetCards struct {
	CurrencySimbol string
	PriceCaption   string

	CSSFlexGap string

	Cards []*WidgetCardVerticalInfo
}

func WidgetCards(params *ParamsWidgetCards) hxprimitives.Node {
	return hxhtml.Div(
		append(
			[]hxprimitives.Node{
				hxprimitives.AttrCSS(
					hxhelpers.Sprintf(
						"display:flex;flex-direction:row;gap:%s;",

						params.CSSFlexGap,
					),
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
	)
}
