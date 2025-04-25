package widgets

import (
	hxhelpers "github.com/TudorHulban/hx-core/helpers"
	hxhtml "github.com/TudorHulban/hx-core/html"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
)

func WidgetSlots(slotsCaptions []string, numberColumns uint8) hxprimitives.Node {
	element := func(caption string) hxprimitives.Node {
		return hxprimitives.Raw(
			hxhelpers.Sprintf(
				`<button class="time-slot" type="button" onclick="handletimeclick('%s')">%s</button>`,

				caption,
				caption,
			),
		)
	}

	rows := []hxprimitives.Node{
		hxprimitives.Raw(
			`<script>
			function handletimeclick(slot){console.log('time slot clicked:', slot);};
			</script>`,
		),
	}

	currentRow := make([]hxprimitives.Node, 0)

	for ix, slot := range slotsCaptions {
		currentRow = append(
			currentRow,
			element(slot),
		)

		if (ix > 0 && (ix+1)%int(numberColumns) == 0) || ix == len(slotsCaptions)-1 {
			rows = append(
				rows,
				hxhtml.Div(
					append(
						[]hxprimitives.Node{
							hxprimitives.AttrClass("hours-row"),
						},

						currentRow...,
					)...,
				),
			)

			currentRow = currentRow[:0]
		}
	}

	return hxhtml.Div(
		append(
			[]hxprimitives.Node{
				hxprimitives.AttrClass("hours-grid"),
			},
			rows...,
		)...,
	)
}
