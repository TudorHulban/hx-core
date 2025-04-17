package hxhtml

import hxprimitives "github.com/TudorHulban/hx-core/primitives"

func A(children ...hxprimitives.Node) hxprimitives.Node {
	return hxprimitives.El(
		"a",
		children...,
	)
}

func Aside(children ...hxprimitives.Node) hxprimitives.Node {
	return hxprimitives.El(
		"aside",
		children...,
	)
}

func Div(children ...hxprimitives.Node) hxprimitives.Writer {
	return hxprimitives.El(
		"div",
		children...,
	)
}

func Label(children ...hxprimitives.Node) hxprimitives.Writer {
	return hxprimitives.El(
		"label",
		children...,
	)
}

func Form(children ...hxprimitives.Node) hxprimitives.Node {
	return hxprimitives.El(
		"form",
		children...,
	)
}

func FormWithID(cssID string, children ...hxprimitives.Node) hxprimitives.Node {
	return hxprimitives.ElWId(
		"form",
		cssID,
		children...,
	)
}

func H1(children ...hxprimitives.Node) hxprimitives.Writer {
	return hxprimitives.El(
		"h1",
		children...,
	)
}

func H2(children ...hxprimitives.Node) hxprimitives.Writer {
	return hxprimitives.El(
		"h2",
		children...,
	)
}

func H3(children ...hxprimitives.Node) hxprimitives.Writer {
	return hxprimitives.El(
		"h3",
		children...,
	)
}

func H4(children ...hxprimitives.Node) hxprimitives.Writer {
	return hxprimitives.El(
		"h4",
		children...,
	)
}

func Img(children ...hxprimitives.Node) hxprimitives.Node {
	return hxprimitives.El(
		"img",
		children...,
	)
}

func Input(children ...hxprimitives.Node) hxprimitives.Writer {
	return hxprimitives.El(
		"input",
		children...,
	)
}

func Li(children ...hxprimitives.Node) hxprimitives.Node {
	return hxprimitives.El(
		"li",
		children...,
	)
}

func Link(children ...hxprimitives.Node) hxprimitives.Node {
	return hxprimitives.El(
		"link",
		children...,
	)
}

func TextArea(children ...hxprimitives.Node) hxprimitives.Writer {
	return hxprimitives.El(
		"textarea",
		children...,
	)
}

func Table(children ...hxprimitives.Node) hxprimitives.Writer {
	return hxprimitives.El(
		"table",
		children...,
	)
}

func THead(children ...hxprimitives.Node) hxprimitives.Node {
	return hxprimitives.El(
		"thead",
		children...,
	)
}

func TBody(children ...hxprimitives.Node) hxprimitives.Writer {
	return hxprimitives.El(
		"tbody",
		children...,
	)
}

func Th(children ...hxprimitives.Node) hxprimitives.Writer {
	return hxprimitives.El(
		"th",
		children...,
	)
}

func Tr(children ...hxprimitives.Node) hxprimitives.Writer {
	return hxprimitives.El(
		"tr",
		children...,
	)
}

func Nav(children ...hxprimitives.Node) hxprimitives.Node {
	return hxprimitives.El(
		"nav",
		children...,
	)
}

func P(children ...hxprimitives.Node) hxprimitives.Node {
	return hxprimitives.El(
		"p",
		children...,
	)
}

func Ul(children ...hxprimitives.Node) hxprimitives.Node {
	return hxprimitives.El(
		"ul",
		children...,
	)
}

func Ol(children ...hxprimitives.Node) hxprimitives.Node {
	return hxprimitives.El(
		"ol",
		children...,
	)
}

func Span(children ...hxprimitives.Node) hxprimitives.Writer {
	return hxprimitives.El(
		"span",
		children...,
	)
}
