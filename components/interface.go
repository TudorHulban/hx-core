package hxcomponents

import "io"

type Node interface {
	isAttribute() bool
	Render(w io.Writer) error
}

var _ Node = attribute{}
var _ Node = Writer(
	func(w io.Writer) error {
		return nil
	},
)
