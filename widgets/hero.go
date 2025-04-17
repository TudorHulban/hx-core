package widgets

import (
	hxcomponents "github.com/TudorHulban/hx-core/components"
	hxhtml "github.com/TudorHulban/hx-core/html"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
)

type ParamsWidgetHero struct {
	Title   string
	Message string

	ButtonPrimaryInfo   hxcomponents.ParamsElementARef
	ButtonSecondaryInfo hxcomponents.ParamsElementARef

	hxcomponents.ParamsImage
}

func WidgetHero(params *ParamsWidgetHero) hxprimitives.Node {
	return hxhtml.Div(
		hxprimitives.AttrClass(
			"hero",
		),

		hxhtml.Div(
			hxprimitives.AttrClass(
				"hero-content",
			),

			hxhtml.H1(
				hxprimitives.Text(
					params.Title,
				),
			),

			hxhtml.P(
				hxprimitives.Text(
					params.Message,
				),
			),

			hxhtml.Div(
				hxprimitives.AttrClass(
					"hero-buttons",
				),

				hxprimitives.Raw(
					hxcomponents.ElementARef(
						&params.ButtonPrimaryInfo,
					),
				),

				hxprimitives.Raw(
					hxcomponents.ElementARef(
						&params.ButtonSecondaryInfo,
					),
				),
			),
		),

		hxhtml.Div(
			hxprimitives.AttrClass(
				"hero-image",
			),

			hxcomponents.Img(
				&params.ParamsImage,
			),
		),
	)
}
