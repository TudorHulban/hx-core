package hxcomponents

import (
	"fmt"

	hxprimitives "github.com/TudorHulban/hx-core/primitives"
)

type ParamsButtonSubmit struct {
	Label    string
	CSSClass string
	CSSID    string

	HXActionType     string
	HXActionEndpoint string

	HXRedirectTo      string
	HXSwapElements    []string // includes CSS ID sanitization
	HXRequireElements []string // includes CSS ID sanitization
	HXSendElements    []string // includes CSS ID sanitization

	HXEnableElements  []string // includes CSS ID sanitization
	HXDisableElements []string // includes CSS ID sanitization

	IsHXUpload bool
	IsDisabled bool
	IsNewTab   bool
}

func ButtonSubmit(params *ParamsButtonSubmit) hxprimitives.Node {
	return hxprimitives.Button(
		hxprimitives.AttrType("submit"),
		hxprimitives.AttrIDLength(params.CSSID),
		hxprimitives.AttrClassLength(params.CSSClass),

		hxprimitives.Text(
			params.Label,
		),

		hxprimitives.If(
			len(params.HXActionType) == 0 || params.IsDisabled,

			hxprimitives.Attr(
				"disabled",
			),
		),

		hxprimitives.If(
			len(params.HXActionType) > 0,

			hxprimitives.AttrWithValue(
				params.HXActionType,
				params.HXActionEndpoint,
			),
		),

		hxprimitives.AttrHXRedirectLength(params.HXRedirectTo),
		hxprimitives.AttrHXSwap(params.HXSwapElements...),
		hxprimitives.AttrHXRequire(params.HXRequireElements...),
		hxprimitives.AttrHXSend(params.HXSendElements...),
		hxprimitives.AttrHXEnable(params.HXEnableElements...),
		hxprimitives.AttrHXDisable(params.HXDisableElements...),
		hxprimitives.AttrHXUpload(params.IsHXUpload),

		hxprimitives.If(
			params.IsNewTab,

			hxprimitives.OnClick(
				fmt.Sprintf(
					`window.open('%s','_blank')`,
					params.HXActionEndpoint,
				),
			),
		),
	)
}
