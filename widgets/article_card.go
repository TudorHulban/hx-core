package widgets

import (
	hxcomponents "github.com/TudorHulban/hx-core/components"
	hxhtml "github.com/TudorHulban/hx-core/html"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
)

type ParamsWidgetArticleCard struct {
	hxcomponents.ParamsImage

	Category string
	Title    string
	MetaInfo string
}

func WidgetArticleCard(params *ParamsWidgetArticleCard) hxprimitives.Node {
	return hxhtml.Article(
		hxprimitives.AttrClass("article-card"),

		hxhtml.Div(
			hxprimitives.AttrClass("article-card-image-container"),

			hxcomponents.Img(
				&params.ParamsImage,
			),
		),

		hxhtml.Div(
			hxprimitives.AttrClass("article-card-content"),

			hxhtml.P(
				hxprimitives.AttrClass("article-card-category"),
				hxprimitives.Text(
					params.Category,
				),
			),

			hxhtml.H2(
				hxprimitives.AttrClass("article-card-title"),

				hxprimitives.Text(
					params.Title,
				),
			),

			hxhtml.P(
				hxprimitives.AttrClass("article-card-meta"),

				hxprimitives.Text(
					params.MetaInfo,
				),
			),
		),
	)
}
