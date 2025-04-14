package hxprimitives

import (
	"io"
)

func Button(children ...Node) Node {
	return El(
		"button",
		children...,
	)
}

func Body(children ...Node) Node {
	return El(
		"body",
		children...,
	)
}

func Doctype(node Node) Node {
	return Writer(
		func(w io.Writer) error {
			if _, errWrite := w.Write([]byte("<!doctype html>")); errWrite != nil {
				return errWrite
			}

			return node.Render(w)
		},
	)
}

func Head(children ...Node) Node {
	return El(
		"head",
		children...,
	)
}

func HTML(children ...Node) Node {
	return El(
		"html",
		children...,
	)
}

func Meta(children ...Node) Node {
	return El(
		"meta",
		children...,
	)
}

func Title(children ...Node) Node {
	return El(
		"title",
		children...,
	)
}

func Script(children ...Node) Node {
	return El(
		"script",
		children...,
	)
}

func SVG(children ...Node) Node {
	return El(
		"svg",
		children...,
	)
}

func Path(children ...Node) Node {
	return El(
		"path",
		children...,
	)
}
