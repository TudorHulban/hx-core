package hxcomponents

import (
	"strings"

	hxhelpers "github.com/TudorHulban/hx-core/helpers"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
)

type InputCheckbox struct {
	CSSDivID    string
	CSSDivClass string

	CSSInputName string
	CSSInputID   string

	ElementName      string
	LabelElementName string

	IsChecked  bool
	IsDisabled bool
}

func (el InputCheckbox) tagCheckbox() string {
	var builder strings.Builder

	builder.WriteString(`type="checkbox" `)
	builder.WriteString(`id="`)
	builder.WriteString(el.CSSInputID)
	builder.WriteString(`" `)
	builder.WriteString(`name="`)
	builder.WriteString(
		hxhelpers.Coalesce(
			el.ElementName,
			el.LabelElementName,
		),
	)
	builder.WriteString(`" `)

	if el.IsChecked {
		builder.WriteString(`checked `)
	}

	if el.IsDisabled {
		builder.WriteString(`disabled `)
	}

	// Trim any trailing space and wrap in <input> tag
	return hxhelpers.Sprintf(
		`<input %s>`,
		strings.TrimSpace(builder.String()),
	)
}

func (el InputCheckbox) tagLabel() string {
	return `<label for="` + el.CSSInputID + `">` + el.LabelElementName + `</label>`
}

func (el InputCheckbox) wrapDiv(content string) string {
	return hxhelpers.Sprintf(
		`<div %s %s>%s</div>`,

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

func (el InputCheckbox) Raw() hxprimitives.Node {
	return hxprimitives.Raw(
		el.wrapDiv(
			strings.Join(
				[]string{
					el.tagCheckbox(),
					el.tagLabel(),
				},
				"\n",
			),
		),
	)
}
