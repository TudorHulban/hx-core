package widgets

import (
	"time"

	hxhelpers "github.com/TudorHulban/hx-core/helpers"
	hxhtml "github.com/TudorHulban/hx-core/html"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
)

type ParamsWidgetInputDate struct {
	DateValue   time.Time
	HowManyDays uint8
}

func WidgetInputDate(params *ParamsWidgetInputDate) hxprimitives.Node {
	return hxhtml.Div(
		hxprimitives.AttrClass("input-date"),

		hxprimitives.Raw(
			hxhelpers.Sprintf(
				`<input type="date" value="%s" min="%s" max="%s"/>`,

				params.DateValue.Format("2006-01-02"),
				params.DateValue.AddDate(0, 0, 1).Format("2006-01-02"),
				params.DateValue.AddDate(0, 0, int(params.HowManyDays)).Format("2006-01-02"),
			),
		),
	)
}
