package hxcomponents

import (
	"strings"

	hxcore "github.com/TudorHulban/hx-core"
	hxhelpers "github.com/TudorHulban/hx-core/helpers"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
)

type InputSelect struct {
	CSSDivID    string
	CSSDivClass string

	CSSInputName string
	CSSInputID   string

	ElementName      string
	LabelElementName string // translated value

	SelectedValue string
	SelectValues  []string
	HX            hxcore.HXInfo

	WithEmptyOption bool
	IsDisabled      bool
}

const optionEmpty = `<option label=" "></option>`

func (el InputSelect) optionSelect(value string) string {
	return hxhelpers.Sprintf(
		`<option value="%s"%s>%s</option>`,

		strings.ToLower(value),
		hxhelpers.Ternary(
			value == el.SelectedValue,

			" selected",
			"",
		),
		value,
	)
}

func (el InputSelect) generateOptions() string {
	options := hxhelpers.ForEachValue(
		el.SelectValues,
		el.optionSelect,
	)

	if el.WithEmptyOption {
		options = append(
			[]string{
				optionEmpty,
			},
			options...,
		)
	}

	return strings.Join(options, "\n")
}

// attributes should be surrounded by space / " ".
func (el InputSelect) selectAttributes() string {
	var sb strings.Builder

	sb.WriteString(
		hxhelpers.Ternary(
			len(el.CSSInputID) > 0,

			`id="`+el.CSSInputID+`"`+" ",
			"",
		),
	)

	sb.WriteString(
		hxhelpers.Ternary(
			len(el.CSSInputName) > 0,

			`name="`+strings.ToLower(el.CSSInputName)+`"`+" ",
			`name="`+strings.ToLower(hxhelpers.Coalesce(el.ElementName, el.LabelElementName))+`"`+" ",
		),
	)

	sb.WriteString(
		hxhelpers.Ternary(
			len(el.HX.Method) > 0,

			el.HX.Method+`="`+el.HX.Endpoint+`"`+" ",
			"",
		),
	)

	sb.WriteString(
		hxhelpers.Ternary(
			len(el.HX.Swaps) > 0,

			hxcore.HXSwap+`="`+hxhelpers.SanitizeCSSIds(el.HX.Swaps)+`"`+" ",
			"",
		),
	)

	sb.WriteString(
		hxhelpers.Ternary(
			len(el.HX.Sends) > 0,

			hxcore.HXSend+`="`+hxhelpers.SanitizeCSSIds(el.HX.Sends)+`"`+" ",
			"",
		),
	)

	sb.WriteString(
		hxhelpers.Ternary(
			len(el.HX.Require) > 0,

			hxcore.HXRequire+`="`+hxhelpers.SanitizeCSSIds(el.HX.Require)+`"`+" ",
			"",
		),
	)

	sb.WriteString(
		hxhelpers.Ternary(
			len(el.HX.OnChangeEnable) > 0,

			hxcore.HXOnChangeEnable+`="`+hxhelpers.SanitizeCSSIds(el.HX.OnChangeEnable)+`"`+" ",
			"",
		),
	)

	sb.WriteString(
		hxhelpers.Ternary(
			el.IsDisabled,

			"disabled"+" ",
			"",
		),
	)

	return sb.String()
}

func (el InputSelect) tagLabel() string {
	return hxhelpers.Ternary(
		len(el.CSSInputID) > 0,

		hxhelpers.Sprintf(
			`<label for="%s">%s:</label>`,

			el.CSSInputID,
			el.LabelElementName,
		),
		hxhelpers.Sprintf(
			`<label>%s:</label>`,

			el.LabelElementName,
		),
	)
}

func (el InputSelect) RawSelect() hxprimitives.Node {
	return hxprimitives.Raw(
		hxhelpers.Sprintf(
			`<select %s>%s</select>`,

			el.selectAttributes(),
			el.generateOptions(),
		),
	)
}

func (el InputSelect) Raw() hxprimitives.Node {
	return hxprimitives.Raw(
		hxhelpers.Sprintf(
			`<div%s%s>%s</div>`,

			hxhelpers.Ternary(
				len(el.CSSDivID) > 0,

				` id="`+el.CSSDivID+`"`,
				"",
			),

			hxhelpers.Ternary(
				len(el.CSSDivClass) > 0,

				` class="`+el.CSSDivClass+`"`,
				"",
			),

			strings.Join(
				[]string{
					el.tagLabel(),
					hxhelpers.Sprintf(
						`<select %s>%s</select>`,

						el.selectAttributes(),
						el.generateOptions(),
					),
				},
				"\n",
			),
		),
	)
}
