package widgets

import (
	hxhtml "github.com/TudorHulban/hx-core/html"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
)

type ParamsWidgetCard struct {
	Title string
}

func WidgetCard(params *ParamsWidgetCard) hxprimitives.Node {
	return hxhtml.Div(
		hxprimitives.AttrClass("has-post-thumbnail"),

		hxhtml.Div(
			hxprimitives.AttrClass("post-wrapper"),

			hxhtml.P(
				hxprimitives.AttrClass("post-thumbnail"),

				hxhtml.A(
					hxprimitives.Href("action1"),

					hxhtml.Img(
						hxprimitives.AttrWithValue("sizes", "auto, (max-width: 920px) 100vw, 920px"),
						hxprimitives.AttrWithValue("decoding", "async"),
						hxprimitives.AttrWithValue("height", "220"),
						hxprimitives.AttrWithValue("alt", ""),
						hxprimitives.AttrWithValue("srcset", "https://themes.getmotopress.com/bro-barbershop/wp-content/uploads/sites/64/2024/01/service-2.jpg 920w, https://themes.getmotopress.com/bro-barbershop/wp-content/uploads/sites/64/2024/01/service-2-300x300.jpg 300w, https://themes.getmotopress.com/bro-barbershop/wp-content/uploads/sites/64/2024/01/service-2-150x150.jpg 150w, https://themes.getmotopress.com/bro-barbershop/wp-content/uploads/sites/64/2024/01/service-2-768x768.jpg 768w"),
						hxprimitives.AttrWithValue("loading", "lazy"),
						hxprimitives.AttrWithValue("width", "220"),
						hxprimitives.AttrWithValue("src", "https://themes.getmotopress.com/bro-barbershop/wp-content/uploads/sites/64/2024/01/service-2.jpg"),
						hxprimitives.AttrClass("attachment-post-thumbnail bw-to-color"),
					),
				),
			),

			hxhtml.Div(
				hxprimitives.AttrClass("post-title-wrapper"),

				hxhtml.H2(
					hxprimitives.AttrClass("entry-title post-title"),

					hxhtml.A(
						hxprimitives.Href("https://themes.getmotopress.com/bro-barbershop/service/washing-head/"),
						hxprimitives.Text(params.Title),
					),
				),
			),
		),
	)
}
