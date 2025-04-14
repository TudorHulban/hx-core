package hxcomponents

import (
	"fmt"
	"strings"

	hxcore "github.com/TudorHulban/hx-core"
	hxhelpers "github.com/TudorHulban/hx-core/helpers"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
)

type InputSimple struct {
	CSSDivID    string
	CSSDivClass string

	CSSInputName string
	CSSInputID   string

	ElementName      string
	LabelElementName string

	TypeInput string // see https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input#input_types
	TextInput string

	HX hxcore.HXInfo

	IsDisabled bool
}

func (el InputSimple) tagInput(toLowerElementName string) string {
	attributes := []string{
		`type="` + el.TypeInput + `"`,

		hxhelpers.Ternary(
			len(el.CSSInputID) > 0,

			`id="`+el.CSSInputID+`"`,
			"",
		),

		`name="` + toLowerElementName + `"`,
	}

	if len(el.TextInput) > 0 && el.TextInput != "0" {
		attributes = append(
			attributes,

			`value="`+el.TextInput+`"`,
		)
	} else {
		if el.HX.LengthMax > 0 {
			attributes = append(
				attributes,

				fmt.Sprintf(
					`%s=%d`,

					hxcore.HXMax,
					el.HX.LengthMax,
				),
			)
		}

		if el.HX.LengthMin > 0 {
			attributes = append(
				attributes,

				fmt.Sprintf(
					`%s=%d`,

					hxcore.HXMin,
					el.HX.LengthMin,
				),
			)
		}

		if el.HX.HavePasswordsValues() {
			attributes = append(
				attributes,

				hxcore.HXValidationPasswords+`="`+
					hxhelpers.SanitizeCSSId(el.HX.CSSIDValidationPasswords[0])+
					`,`+
					hxhelpers.SanitizeCSSId(el.HX.CSSIDValidationPasswords[1])+
					`"`,
			)
		}

		if len(el.HX.CSSIDValidationDisable) > 0 {
			attributes = append(
				attributes,

				hxcore.HXValidationDisable+`="`+el.HX.CSSIDValidationDisable+`"`,
			)
		}
	}

	if el.IsDisabled {
		attributes = append(
			attributes,

			"disabled",
		)
	}

	return fmt.Sprintf(
		`<input %s>`,

		strings.Join(attributes, " "),
	)
}

func (el InputSimple) wrapDiv(content string) string {
	return hxhelpers.Sprintf(
		`<div %s%s>%s</div>`,

		hxhelpers.Ternary(
			len(el.CSSDivID) > 0,

			`id="`+el.CSSDivID+`"`,
			"",
		),
		hxhelpers.Ternary(
			len(el.CSSDivClass) > 0,

			`class="`+el.CSSDivClass+`"`,
			"",
		),
		content,
	)
}

func (el InputSimple) RawSelect() hxprimitives.Node {
	return hxprimitives.Raw(
		el.tagInput(
			strings.ToLower(
				hxhelpers.Coalesce(
					el.ElementName,
					el.LabelElementName,
				),
			),
		),
	)
}

func (el InputSimple) Raw() hxprimitives.Node {
	return hxprimitives.Raw(
		el.wrapDiv(
			strings.Join(
				[]string{
					hxhelpers.Ternary(
						len(el.CSSInputID) > 0,

						`<label for="`+el.CSSInputID+`">`+el.LabelElementName+`:</label>`,
						`<label>`+el.LabelElementName+`:</label>`,
					),

					el.tagInput(
						strings.ToLower(
							hxhelpers.Coalesce(
								el.ElementName,
								el.LabelElementName,
							),
						),
					),
				},
				"",
			),
		),
	)
}
