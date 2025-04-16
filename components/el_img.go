package hxcomponents

import (
	hxhelpers "github.com/TudorHulban/hx-core/helpers"
	hxhtml "github.com/TudorHulban/hx-core/html"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
)

type ParamsImage struct {
	ImageSquareSize string
	ImageSource     string
	ImageAlt        string
}

func Img(params *ParamsImage) hxprimitives.Node {
	return hxhtml.Img(
		hxprimitives.AttrWithValue(
			"sizes",

			hxhelpers.Sprintf(
				"auto, (max-width: %s) 100vw, %s",

				params.ImageSquareSize,
				params.ImageSquareSize,
			),
		),
		hxprimitives.AttrCSS("object-fit:cover;"),
		hxprimitives.AttrWithValue("decoding", "async"),
		hxprimitives.AttrWithValue("height", params.ImageSquareSize),
		hxprimitives.AttrWithValue("width", params.ImageSquareSize),
		hxprimitives.AttrWithValue("alt", params.ImageAlt),
		hxprimitives.AttrWithValue("loading", "lazy"),
		hxprimitives.AttrWithValue("src", params.ImageSource),
		hxprimitives.AttrClass("attachment-post-thumbnail"),
	)
}
