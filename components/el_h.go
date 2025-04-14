package hxcomponents

import (
	hxhtml "github.com/TudorHulban/hx-core/html"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
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

func H4Raw(params *ParamsH) string {
	return `<h4 id="` + params.CSSID + `">` + params.Text + `</h4>`
}

func H5Raw(params *ParamsH) string {
	return `<h5 id="` + params.CSSID + `">` + params.Text + `</h5>`
}

func H6Raw(params *ParamsH) string {
	return `<h6 id="` + params.CSSID + `">` + params.Text + `</h6>`
}

func H3(params *ParamsH) hxprimitives.Node {
	return hxhtml.H3(
		hxprimitives.AttrIDLength(
			params.CSSID,
		),

		hxprimitives.Text(
			params.Text,
		),
	)
}
