package hxcomponents

import hxprimitives "github.com/TudorHulban/hx-core/primitives"

type Page struct {
	Title       string
	Description string
	Language    string

	Head []hxprimitives.Node
	Body []hxprimitives.Node
}

func (p *Page) Build() hxprimitives.Node {
	return hxprimitives.Doctype(
		hxprimitives.HTML(
			hxprimitives.Lang(p.Language),

			hxprimitives.Head(
				append(
					[]hxprimitives.Node{
						hxprimitives.Meta(
							hxprimitives.Charset("utf-8"),
						),
						hxprimitives.Meta(
							hxprimitives.Name("viewport"),
							hxprimitives.Content("width=device-width, initial-scale=1"),
						),
						hxprimitives.Title(
							hxprimitives.Text(p.Title),
						),
						hxprimitives.If(
							len(p.Description) > 0,
							hxprimitives.Meta(
								hxprimitives.Name("description"),
								hxprimitives.Content(p.Description),
							),
						),
					},
					p.Head...,
				)...,
			),

			hxprimitives.Body(
				p.Body...,
			),
		),
	)
}
