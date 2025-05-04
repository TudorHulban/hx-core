package hxcomponents

import (
	hxhtml "github.com/TudorHulban/hx-core/html"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
)

type ParamsNewFormThreeContainers struct {
	SymbolEntry   string // form logo symbol
	TextForm      string
	CSSIDTextForm string

	IDDivEnclosing    string
	IDForm            string
	IDDivContainers   string // link media queries to this ID.
	IDContainerLeft   string
	IDContainerMiddle string
	IDContainerRight  string

	ElementsInputLeft   []hxprimitives.Node
	ElementsInputMiddle []hxprimitives.Node
	ElementsInputRight  []hxprimitives.Node

	Buttons []hxprimitives.Node
}

func NewFormThreeContainers(params *ParamsNewFormThreeContainers) hxprimitives.Node {
	return hxhtml.Div(
		hxprimitives.AttrIDLength(params.IDDivEnclosing),

		hxprimitives.If(
			len(params.SymbolEntry) > 0 || len(params.TextForm) > 0,

			hxhtml.Div(
				hxprimitives.AttrCSS(
					"display: flex;",
				),

				hxprimitives.If(
					len(params.SymbolEntry) > 0,

					hxhtml.Span(
						hxprimitives.AttrClass("material-symbols-outlined"),
						hxprimitives.Text(
							params.SymbolEntry,
						),
					),
				),

				hxprimitives.If(
					len(params.TextForm) > 0,

					H3(
						&ParamsH{
							Text:  params.TextForm,
							CSSID: params.CSSIDTextForm,
						},
					),
				),
			),
		),

		hxhtml.FormWithID(
			params.IDForm,

			append(
				[]hxprimitives.Node{
					hxprimitives.AutocompleteOff(),

					hxhtml.Div(
						[]hxprimitives.Node{
							hxprimitives.AttrID(params.IDDivContainers),

							RenderFormContainer(
								&ParamsRenderFormContainer{
									IDContainer: params.IDContainerLeft,
									Elements:    params.ElementsInputLeft,
								},
							),

							RenderFormContainer(
								&ParamsRenderFormContainer{
									IDContainer: params.IDContainerMiddle,
									Elements:    params.ElementsInputMiddle,
								},
							),

							RenderFormContainer(
								&ParamsRenderFormContainer{
									IDContainer: params.IDContainerRight,
									Elements:    params.ElementsInputRight,
								},
							),
						}...,
					),
				},

				hxprimitives.If(
					len(params.Buttons) > 0,

					hxhtml.Div(
						append(
							[]hxprimitives.Node{
								hxprimitives.AttrCSS(`display:flex;width:100%;flex-wrap:nowrap;gap:10px;`),
							},
							params.Buttons...,
						)...,
					),
				),
			)...,
		),
	)
}
