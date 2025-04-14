package hxcomponents

import (
	hxcore "github.com/TudorHulban/hx-core"
	hxhelpers "github.com/TudorHulban/hx-core/helpers"
)

type ParamsElementARef struct {
	Route    string
	TicketID string
	Label    string

	TargetsMultiswap string
	TargetsSend      string
}

func ElementARef(params *ParamsElementARef) string {
	return hxhelpers.Sprintf(
		`<a href="#" %s="%s/%s" %s="%s" %s="%s">%s</a>`,

		hxcore.HXPOST,
		params.Route,
		params.TicketID,

		hxcore.HXSwap,
		params.TargetsMultiswap,

		hxcore.HXSend,
		params.TargetsSend,

		params.Label,
	)
}
