package hxcomponents

import (
	"strings"

	hxcore "github.com/TudorHulban/hx-core"
	hxhelpers "github.com/TudorHulban/hx-core/helpers"
)

func AttrHXRedirect(value string) attribute {
	return AttrWithValue(
		hxcore.HXRedirect,
		value,
	)
}

func AttrHXRedirectLength(value string) attribute {
	if len(value) > 0 {
		return AttrWithValue(
			hxcore.HXRedirect,
			value,
		)
	}

	return attribute{}
}

func AttrHXRequire(values ...string) attribute {
	return AttrWithValueIf(
		len(values) > 0,

		hxcore.HXRequire,
		strings.Join(
			hxhelpers.SanitizeCSSIdsSlice(values...),
			",",
		),
	)
}

func AttrHXSwap(values ...string) attribute {
	return AttrWithValueIf(
		len(values) > 0,

		hxcore.HXSwap,
		strings.Join(
			hxhelpers.SanitizeCSSIdsSlice(values...),
			",",
		),
	)
}

func AttrHXSend(values ...string) attribute {
	return AttrWithValueIf(
		len(values) > 0,

		hxcore.HXSend,
		strings.Join(
			hxhelpers.SanitizeCSSIdsSlice(values...),
			",",
		),
	)
}

func AttrHXEnable(values ...string) attribute {
	return AttrWithValueIf(
		len(values) > 0,

		hxcore.HXDirectEnable,
		strings.Join(
			hxhelpers.SanitizeCSSIdsSlice(values...),
			",",
		),
	)
}

func AttrHXDisable(values ...string) attribute {
	return AttrWithValueIf(
		len(values) > 0,

		hxcore.HXDirectDisable,
		strings.Join(
			hxhelpers.SanitizeCSSIdsSlice(values...),
			",",
		),
	)
}

func AttrHXUpload(withUpload bool) attribute {
	if withUpload {
		return Attr(
			hxcore.HXUpload,
		)
	}

	return attribute{}
}

func AttrHXPOST(value string) attribute {
	return AttrWithValue(
		hxcore.HXPOST,
		value,
	)
}

func AttrHXShow(values ...string) attribute {
	return AttrWithValueIf(
		len(values) > 0,

		hxcore.HXShow,
		strings.Join(
			values,
			",",
		),
	)
}

func AttrHXHide(values ...string) attribute {
	return AttrWithValueIf(
		len(values) > 0,

		hxcore.HXHide,
		strings.Join(
			values,
			",",
		),
	)
}
