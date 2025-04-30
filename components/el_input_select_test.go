package hxcomponents

import (
	"fmt"
	"os"
	"testing"

	hxcore "github.com/TudorHulban/hx-core"
)

func TestElementSelectInput(t *testing.T) {
	el := InputSelect{
		CSSDivClass:      "class-div",
		CSSInputID:       "id-input",
		LabelElementName: "label",

		HX: hxcore.HXInfo{
			Swaps: []string{
				"id1",
			},
		},

		SelectValues: []string{
			"a",
			"b",
		},
	}

	// <div  class="class-div"><label for="id-input">label:</label>
	// <select id="id-input" name="label" hx-swap="#id1" ><option value="a">a</option>
	// <option value="b">b</option></select>
	// </div>

	fmt.Println(
		el.
			Raw().
			Render(os.Stdout),
	)
}

func TestEmptyElementSelectInput(t *testing.T) {
	el := InputSelect{
		CSSDivClass:      "class-div",
		CSSInputID:       "id-input",
		LabelElementName: "label",

		HX: hxcore.HXInfo{
			Sends: []string{
				"id1",
			},
		},
	}

	// <div  class="class-div"><label for="id-input">label:</label>
	// <select id="id-input" name="label" hx-send="#id1" ></select>
	// </div>

	fmt.Println(
		el.
			Raw().
			Render(os.Stdout),
	)
}

func TestOnChangeSelectInput(t *testing.T) {
	el := InputSelect{
		CSSDivClass:      "class-div",
		CSSInputID:       "id-input",
		LabelElementName: "label",

		HX: hxcore.HXInfo{
			OnChangeEnable: []string{
				"id1",
			},
		},

		SelectValues: []string{
			"a",
			"b",
		},
	}

	// <div  class="class-div"><label for="id-input">label:</label>
	// <select id="id-input" name="label" hx-swap="#id1" ><option value="a">a</option>
	// <option value="b">b</option></select>
	// </div>

	fmt.Println(
		el.
			Raw().
			Render(os.Stdout),
	)
}
