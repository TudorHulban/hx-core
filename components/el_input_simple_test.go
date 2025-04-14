package hxcomponents

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	hxcore "github.com/TudorHulban/hx-core"
	"github.com/stretchr/testify/require"
)

func TestNoIDElementSimpleInput(t *testing.T) {
	el := InputSimple{
		CSSDivClass:      "class-div",
		LabelElementName: "label",

		TypeInput: "text",
		// TextInput: "some text to show",

		HX: hxcore.HXInfo{
			LengthMax: 50,
		},
	}

	// 	<div  class="class-div">`<label>`+el.LabelElementName+`:</label>`
	// <input type="text"  name="label" hx-max=50></div>

	fmt.Println(
		el.
			Raw().
			Render(os.Stdout),
	)

	expectedOutput := fmt.Sprintf(`<div class="class-div"><label>%s:</label><input type="text"  name="label" hx-max=50></div>`, el.LabelElementName)

	var buf bytes.Buffer
	err := el.Raw().Render(&buf)
	require.NoError(t, err, "render error")

	actualOutput := buf.String()
	require.Equal(t, expectedOutput, actualOutput, "unexpected output")
}
