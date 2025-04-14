package hxprimitives

import (
	"fmt"

	hxcomponents "github.com/TudorHulban/hx-core/components"
	hxhtml "github.com/TudorHulban/hx-core/html"
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

func ButtonSubmit(params *ParamsButtonSubmit) hxcomponents.Node {
	return hxhtml.Button(
		hxcomponents.AttrType("submit"),
		hxcomponents.AttrIDLength(params.CSSID),
		hxcomponents.AttrClassLength(params.CSSClass),

		hxcomponents.Text(
			params.Label,
		),

		hxcomponents.If(
			len(params.HXActionType) == 0 || params.IsDisabled,

			hxcomponents.Attr(
				"disabled",
			),
		),

		hxcomponents.If(
			len(params.HXActionType) > 0,

			hxcomponents.AttrWithValue(
				params.HXActionType,
				params.HXActionEndpoint,
			),
		),

		hxcomponents.AttrHXRedirectLength(params.HXRedirectTo),
		hxcomponents.AttrHXSwap(params.HXSwapElements...),
		hxcomponents.AttrHXRequire(params.HXRequireElements...),
		hxcomponents.AttrHXSend(params.HXSendElements...),
		hxcomponents.AttrHXEnable(params.HXEnableElements...),
		hxcomponents.AttrHXDisable(params.HXDisableElements...),
		hxcomponents.AttrHXUpload(params.IsHXUpload),

		hxcomponents.If(
			params.IsNewTab,

			hxcomponents.OnClick(
				fmt.Sprintf(
					`window.open('%s','_blank')`,
					params.HXActionEndpoint,
				),
			),
		),
	)
}
