package hxprimitives

import (
	"io"
	"text/template"
)

func Text(text string) Writer {
	return func(w io.Writer) error {
		_, errWrite := w.Write(
			[]byte(template.HTMLEscapeString(text)),
		)

		return errWrite
	}
}

func Raw(text string) Writer {
	return func(w io.Writer) error {
		_, errWrite := w.Write(
			[]byte(text),
		)

		return errWrite
	}
}

func If(condition bool, node Node) Node {
	if condition {
		return node
	}

	return nil
}

func Ternary(condition bool, node1, node2 Node) Node {
	if condition {
		return node1
	}

	return node2
}
