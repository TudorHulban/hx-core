package hxhtml

import hxprimitives "github.com/TudorHulban/hx-core/primitives"

func Article(children ...hxprimitives.Node) hxprimitives.Node {
	return hxprimitives.El(
		"article",
		children...,
	)
}
