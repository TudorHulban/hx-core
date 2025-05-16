package pagecss

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNormalizeCSS(t *testing.T) {
	t.Run(
		"1. usual CSS",
		func(t *testing.T) {
			normalized, errNormalization := NormalizeCSS(
				&ParamsNormalizeCSS{
					CSS: `
	h1 {
font-size: 75px;
line-height: 1.15;
}

h2 {
font-size: 45px;
}
	`,
					CSSNumberSpaces:       5,
					IncrementNumberSpaces: 2,
				},
			)
			require.NoError(t, errNormalization)

			fmt.Println(normalized)
		},
	)

	t.Run(
		"2. nested CSS",
		func(t *testing.T) {
			normalized, errNormalization := NormalizeCSS(
				&ParamsNormalizeCSS{
					CSS: `
	.parent {
     .child {
    property: value;
  }
}
	`,
					CSSNumberSpaces:       5,
					IncrementNumberSpaces: 2,
				},
			)

			require.NoError(t, errNormalization)
			fmt.Println(normalized)
		},
	)

	t.Run(
		"3. CSS with comments",
		func(t *testing.T) {
			normalized, errNormalization := NormalizeCSS(
				&ParamsNormalizeCSS{
					CSS: `
	/* Input CSS */
.class {
  /* This comment contains a closing brace } that confuses the logic */
  property: value;
}

/* Output using the provided logic */
.class {
/* This comment contains a closing brace } that confuses the logic */
property: value; /* This line might not be indented */
}
	`,
					CSSNumberSpaces:       5,
					IncrementNumberSpaces: 2,
				},
			)

			require.NoError(t, errNormalization)
			fmt.Println(normalized)
		},
	)
}
