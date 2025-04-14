package hxprimitives

import (
	hxcomponents "github.com/TudorHulban/hx-core/components"
	hxhtml "github.com/TudorHulban/hx-core/html"
)

type ParamsH struct {
	CSSID string
	Text  string
}

func H1Raw(params *ParamsH) string {
	return `<h1 id="` + params.CSSID + `">` + params.Text + `</h1>`
}

func H2Raw(params *ParamsH) string {
	return `<h2 id="` + params.CSSID + `">` + params.Text + `</h2>`
}

func H3Raw(params *ParamsH) string {
	return `<h3 id="` + params.CSSID + `">` + params.Text + `</h3>`
}

func H5Raw(params *ParamsH) string {
	return `<h5 id="` + params.CSSID + `">` + params.Text + `</h5>`
}

func H6Raw(params *ParamsH) string {
	return `<h6 id="` + params.CSSID + `">` + params.Text + `</h6>`
}

func H3(params *ParamsH) hxcomponents.Node {
	return hxhtml.H3(
		hxcomponents.AttrIDLength(
			params.CSSID,
		),

		hxcomponents.Text(
			params.Text,
		),
	)
}
