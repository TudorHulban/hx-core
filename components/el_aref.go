package hxcomponents

import (
	hxcore "github.com/TudorHulban/hx-core"
	hxhelpers "github.com/TudorHulban/hx-core/helpers"
)

type ParamsElementARef struct {
	CSSClass string

	Route   string
	ItemID  string
	Caption string

	TargetsMultiswap string
	TargetsSend      string
}

func ElementARef(params *ParamsElementARef) string {
	return hxhelpers.Sprintf(
		`<a%shref="#" %s="%s" %s="%s"%s>%s</a>`,

		hxhelpers.Ternary(
			len(params.CSSClass) > 0,

			hxhelpers.Sprintf(
				` class="%s" `,
				params.CSSClass,
			),
			"",
		),

		hxcore.HXPOST,

		hxhelpers.Ternary(
			len(params.ItemID) > 0,

			params.Route+"/"+params.ItemID,
			params.Route,
		),

		// no ternary for swap as not probable.
		hxcore.HXSwap,
		params.TargetsMultiswap,

		hxhelpers.Ternary(
			len(params.TargetsSend) > 0,

			hxhelpers.Sprintf(
				` %s="%s"`,
				hxcore.HXSend,
				params.TargetsSend,
			),
			"",
		),

		params.Caption,
	)
}
