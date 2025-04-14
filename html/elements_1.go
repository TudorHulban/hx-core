package hxhtml

import hxcomponents "github.com/TudorHulban/hx-core/components"

func A(children ...hxcomponents.Node) hxcomponents.Node {
	return hxcomponents.El(
		"a",
		children...,
	)
}

func Aside(children ...hxcomponents.Node) hxcomponents.Node {
	return hxcomponents.El(
		"aside",
		children...,
	)
}

func Div(children ...hxcomponents.Node) hxcomponents.Writer {
	return hxcomponents.El(
		"div",
		children...,
	)
}

func Label(children ...hxcomponents.Node) hxcomponents.Writer {
	return hxcomponents.El(
		"label",
		children...,
	)
}

func Form(children ...hxcomponents.Node) hxcomponents.Node {
	return hxcomponents.El(
		"form",
		children...,
	)
}

func FormWithID(cssID string, children ...hxcomponents.Node) hxcomponents.Node {
	return hxcomponents.ElWId(
		"form",
		cssID,
		children...,
	)
}

func H2(children ...hxcomponents.Node) hxcomponents.Writer {
	return hxcomponents.El(
		"h2",
		children...,
	)
}

func H3(children ...hxcomponents.Node) hxcomponents.Writer {
	return hxcomponents.El(
		"h3",
		children...,
	)
}

func H4(children ...hxcomponents.Node) hxcomponents.Writer {
	return hxcomponents.El(
		"h4",
		children...,
	)
}

func Img(children ...hxcomponents.Node) hxcomponents.Node {
	return hxcomponents.El(
		"img",
		children...,
	)
}

func Input(children ...hxcomponents.Node) hxcomponents.Writer {
	return hxcomponents.El(
		"input",
		children...,
	)
}

func Li(children ...hxcomponents.Node) hxcomponents.Node {
	return hxcomponents.El(
		"li",
		children...,
	)
}

func Link(children ...hxcomponents.Node) hxcomponents.Node {
	return hxcomponents.El(
		"link",
		children...,
	)
}

func TextArea(children ...hxcomponents.Node) hxcomponents.Writer {
	return hxcomponents.El(
		"textarea",
		children...,
	)
}

func Table(children ...hxcomponents.Node) hxcomponents.Writer {
	return hxcomponents.El(
		"table",
		children...,
	)
}

func THead(children ...hxcomponents.Node) hxcomponents.Node {
	return hxcomponents.El(
		"thead",
		children...,
	)
}

func TBody(children ...hxcomponents.Node) hxcomponents.Writer {
	return hxcomponents.El(
		"tbody",
		children...,
	)
}

func Th(children ...hxcomponents.Node) hxcomponents.Writer {
	return hxcomponents.El(
		"th",
		children...,
	)
}

func Tr(children ...hxcomponents.Node) hxcomponents.Writer {
	return hxcomponents.El(
		"tr",
		children...,
	)
}

func Nav(children ...hxcomponents.Node) hxcomponents.Node {
	return hxcomponents.El(
		"nav",
		children...,
	)
}

func P(children ...hxcomponents.Node) hxcomponents.Node {
	return hxcomponents.El(
		"p",
		children...,
	)
}

func Ul(children ...hxcomponents.Node) hxcomponents.Node {
	return hxcomponents.El(
		"ul",
		children...,
	)
}

func Ol(children ...hxcomponents.Node) hxcomponents.Node {
	return hxcomponents.El(
		"ol",
		children...,
	)
}

func Span(children ...hxcomponents.Node) hxcomponents.Writer {
	return hxcomponents.El(
		"span",
		children...,
	)
}
