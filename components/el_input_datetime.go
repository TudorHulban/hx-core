package hxcomponents

import (
	"database/sql"
	"strings"

	hxhelpers "github.com/TudorHulban/hx-core/helpers"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
	hxrequest "github.com/TudorHulban/hx-request"
)

type InputDateTime struct {
	CSSDivClass string
	CSSInputID  string

	ElementName      string
	LabelElementName string

	AfterTime sql.NullTime

	IsDisabled bool
}

func (el InputDateTime) Raw() hxprimitives.Node {
	result := [4]string{
		hxhelpers.Ternary(
			len(el.CSSDivClass) == 0,

			`<div>`,
			hxhelpers.Sprintf(
				`<div class="%s">`,
				el.CSSDivClass,
			),
		),

		hxhelpers.Sprintf(
			`<label for="%s">%s:</label>`,

			el.ElementName,
			el.LabelElementName,
		),

		hxhelpers.Sprintf(
			`<input type="datetime-local" id="%s" name="%s" %s`,

			el.CSSInputID,
			hxhelpers.Coalesce(
				el.ElementName,
				el.LabelElementName,
			),

			hxhelpers.Ternary(
				el.AfterTime.Valid,

				hxhelpers.Sprintf(
					`value="%s"`,
					hxrequest.TimeForInputDateTime(
						el.AfterTime.Time,
						hxrequest.LayoutDateTime,
					),
				),
				"",
			),
		),

		hxhelpers.Ternary(
			el.IsDisabled,

			"disabled></div>",
			"></div>",
		),
	}

	return hxprimitives.Raw(
		strings.Join(result[:], "\n"),
	)
}
