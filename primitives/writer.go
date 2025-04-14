package hxprimitives

import "io"

type Writer func(io.Writer) error

func (wHX Writer) isAttribute() bool {
	return false
}

func (wHX Writer) Render(w io.Writer) error {
	return wHX(w)
}
