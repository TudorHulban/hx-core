package widgets

import (
	"fmt"
	"time"

	hxcomponents "github.com/TudorHulban/hx-core/components"
	hxhtml "github.com/TudorHulban/hx-core/html"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
)

type ArticleMeta struct {
	LastUpdate     time.Time
	Author         string
	NumberComments uint16
}

func (meta ArticleMeta) String() string {
	return fmt.Sprintf(
		`%s | by %s | %d Comments`,

		meta.LastUpdate.Format("January 2, 2006"),
		meta.Author,
		meta.NumberComments,
	)
}

type ParamsWidgetArticleCard struct {
	hxcomponents.ParamsImage
	hxcomponents.ParamsButtonSubmit

	Category       string
	Title          string
	ArticleExcerpt string

	ArticleMeta
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
					params.ArticleMeta.String(),
				),
			),

			hxhtml.P(
				hxprimitives.AttrClass("article-card-excerpt"),

				hxprimitives.Text(
					params.ArticleExcerpt,
				),
			),

			hxhtml.Div(
				hxprimitives.AttrClass("article-card-action"),

				hxcomponents.ButtonSubmit(
					&params.ParamsButtonSubmit,
				),
			),
		),
	)
}

func CSSArticleCard() {}
