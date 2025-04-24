package widgets

import (
	"time"

	hxhelpers "github.com/TudorHulban/hx-core/helpers"
	hxhtml "github.com/TudorHulban/hx-core/html"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
)

type ParamsWidgetInputDate struct {
	CSSID string `valid:"required"`

	DateValue   time.Time
	HowManyDays uint8
}

type ResponseWidgetInputDate struct {
	HTML           hxprimitives.Node
	LinkJavascript hxprimitives.Node
}

func WidgetInputDate(params *ParamsWidgetInputDate) *ResponseWidgetInputDate {
	return &ResponseWidgetInputDate{
		HTML: hxhtml.Div(
			hxprimitives.AttrClass("input-date"),

			hxprimitives.Raw(
				hxhelpers.Sprintf(
					`<input class="flatpickr-input" type="text" id="%s" value="%s" min="%s" max="%s"/>`,

					params.CSSID,
					params.DateValue.Format("2006-01-02"),
					params.DateValue.AddDate(0, 0, 1).Format("2006-01-02"),
					params.DateValue.AddDate(0, 0, int(params.HowManyDays)).Format("2006-01-02"),
				),
			),

			hxprimitives.Raw(
				hxhelpers.Sprintf(
					`<script>
					document.addEventListener('DOMContentLoaded', function() {
					const inputElement = document.getElementById('%s');
					if (inputElement) {
							flatpickr(inputElement, {
								dateFormat: "Y-m-d",
								inline: true,
								defaultDate: inputElement.value,
								minDate: inputElement.getAttribute('min'),
								maxDate: inputElement.getAttribute('max')
							});
						}
					});
					</script>`,

					params.CSSID,
				),
			),
		),

		LinkJavascript: hxprimitives.Raw(
			`<script src="https://cdn.jsdelivr.net/npm/flatpickr"></script>`,
		),
	}
}
