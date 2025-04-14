package hxhtml

import (
	"io"

	hxcomponents "github.com/TudorHulban/hx-core/components"
)

func Button(children ...hxcomponents.Node) hxcomponents.Node {
	return hxcomponents.El(
		"button",
		children...,
	)
}

func Body(children ...hxcomponents.Node) hxcomponents.Node {
	return hxcomponents.El(
		"body",
		children...,
	)
}

func Doctype(node hxcomponents.Node) hxcomponents.Node {
	return hxcomponents.Writer(
		func(w io.Writer) error {
			if _, errWrite := w.Write([]byte("<!doctype html>")); errWrite != nil {
				return errWrite
			}

			return node.Render(w)
		},
	)
}

func Head(children ...hxcomponents.Node) hxcomponents.Node {
	return hxcomponents.El(
		"head",
		children...,
	)
}

func HTML(children ...hxcomponents.Node) hxcomponents.Node {
	return hxcomponents.El(
		"html",
		children...,
	)
}

func Meta(children ...hxcomponents.Node) hxcomponents.Node {
	return hxcomponents.El(
		"meta",
		children...,
	)
}

func Title(children ...hxcomponents.Node) hxcomponents.Node {
	return hxcomponents.El(
		"title",
		children...,
	)
}

func Script(children ...hxcomponents.Node) hxcomponents.Node {
	return hxcomponents.El(
		"script",
		children...,
	)
}

func SVG(children ...hxcomponents.Node) hxcomponents.Node {
	return hxcomponents.El(
		"svg",
		children...,
	)
}

func Path(children ...hxcomponents.Node) hxcomponents.Node {
	return hxcomponents.El(
		"path",
		children...,
	)
}
