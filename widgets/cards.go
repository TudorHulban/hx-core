package widgets

import (
	"fmt"

	hxhelpers "github.com/TudorHulban/hx-core/helpers"
	hxhtml "github.com/TudorHulban/hx-core/html"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
)

type ParamsWidgetCards struct {
	CurrencySimbol string
	PriceCaption   string

	CSSFlexGap string

	Cards []*WidgetCardInfo
}

func WidgetCards(params *ParamsWidgetCards) hxprimitives.Node {
	return hxhtml.Div(
		append(
			[]hxprimitives.Node{
				hxprimitives.AttrCSS(
					fmt.Sprintf(
						"display:flex;flex-direction:row;gap:%s;",

						params.CSSFlexGap,
					),
				),
			},

			hxhelpers.ForEachValue(
				params.Cards,

				func(info *WidgetCardInfo) hxprimitives.Node {
					return WidgetCard(
						&ParamsWidgetCard{
							WidgetCardInfo: *info,

							CurrencySimbol: params.CurrencySimbol,
							PriceCaption:   params.PriceCaption,
						},
					)
				},
			)...,
		)...,
	)
}
