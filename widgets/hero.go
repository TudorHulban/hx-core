package widgets

import (
	hxcomponents "github.com/TudorHulban/hx-core/components"
	hxhtml "github.com/TudorHulban/hx-core/html"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
)

type ParamsWidgetHero struct {
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
					"Expert cuts and classic styles",
				),
			),

			hxhtml.P(
				hxprimitives.Text(
					"experience the finest grooming services in town. from traditional haircuts to modern styling, we've got you covered.",
				),
			),

			hxhtml.Div(
				hxprimitives.AttrClass(
					"hero-buttons",
				),

				hxhtml.A(
					hxprimitives.AttrClass(
						"button primary",
					),

					hxprimitives.Text(
						"Book appointment",
					),
				),

				hxhtml.A(
					hxprimitives.AttrClass(
						"button secondary",
					),

					hxprimitives.Text(
						"View services",
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
