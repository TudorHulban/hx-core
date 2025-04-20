package widgets

import (
	"fmt"
	"testing"
	"time"

	hxcore "github.com/TudorHulban/hx-core"
	hxcomponents "github.com/TudorHulban/hx-core/components"
	hxhtml "github.com/TudorHulban/hx-core/html"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
	"github.com/stretchr/testify/require"
)

func TestWidgetArticleCard(t *testing.T) {
	fragment := WidgetArticleCard(
		&ParamsWidgetArticleCard{
			Category:       "category-g1",
			Title:          "Article-title",
			ArticleExcerpt: "lorem ipsum dolor sit amet, consectetur adipiscing elit. sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",

			ArticleMeta: ArticleMeta{
				LastUpdate:     time.Now(),
				Author:         "John Smith",
				NumberComments: 1,
			},

			ParamsImage: hxcomponents.ParamsImage{
				ImageSquareSize: "260",
				ImageSource:     "https://images.pexels.com/photos/539694/pexels-photo-539694.jpeg",
				ImageAlt:        "1",
			},

			ParamsButtonSubmit: hxcomponents.ParamsButtonSubmit{
				Label: "Continue reading",

				HXActionType: hxcore.HXPOST,
			},
		},
	)

	writer, errWriter := getFileWriter(t.Name() + ".html")
	require.NoError(t, errWriter)

	defer writer.Close()

	page := hxcomponents.Page{
		Head: []hxprimitives.Node{
			hxhtml.Link(
				hxprimitives.Rel("stylesheet"),
				hxprimitives.Href("css_base.css"),
			),
			hxhtml.Link(
				hxprimitives.Rel("stylesheet"),
				hxprimitives.Href("css_site.css"),
			),
			hxhtml.Link(
				hxprimitives.Rel("stylesheet"),
				hxprimitives.Href("article_card.css"),
			),
		},

		Body: []hxprimitives.Node{
			fragment,
		},
	}

	fmt.Println(
		page.Build().Render(writer),
	)
}
