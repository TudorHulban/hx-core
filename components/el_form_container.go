package hxcomponents

import (
	hxhtml "github.com/TudorHulban/hx-core/html"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
)

type ParamsRenderFormContainer struct {
	IDContainer string
	Elements    []hxprimitives.Node
}

func RenderFormContainer(params *ParamsRenderFormContainer) hxprimitives.Node {
	return hxhtml.Div(
		append(
			params.Elements,

			hxprimitives.If(
				len(params.IDContainer) > 0,

				hxprimitives.AttrID(params.IDContainer),
			),
		)...,
	)
}
