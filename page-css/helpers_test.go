package pagecss

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSpecialCSS(t *testing.T) {
	t.Run(
		"1. moz-focus-inner",
		func(t *testing.T) {
			normalized, errNormalization := BeautifyCSS(
				&ParamsUpdateCSS{
					CSS: `
button::-moz-focus-inner, .mpa-employees-list-shortcode .entry-link::-moz-focus-inner,
.wp-block-getwid-post-carousel .entry-link::-moz-focus-inner,
[type="button"]::-moz-focus-inner,
[type="reset"]::-moz-focus-inner,
[type="submit"]::-moz-focus-inner {
	border-style: none;
	padding: 0;
}
	`,
					ParamsSpaces: ParamsSpaces{
						NumberSpaces:              5,
						IncrementWithNumberSpaces: 2,
					},
				},
			)
			require.NoError(t, errNormalization)

			fmt.Println(normalized)
		},
	)
}

func TestBeautifyCSS(t *testing.T) {
	t.Run(
		"1. simple CSS",
		func(t *testing.T) {
			normalized, errNormalization := BeautifyCSS(
				&ParamsUpdateCSS{
					CSS: `
	h1 {
font-size: 75px;
line-height: 1.15;
}

h2 {
font-size: 45px;
}
	`,
					ParamsSpaces: ParamsSpaces{
						NumberSpaces:              5,
						IncrementWithNumberSpaces: 2,
					},
				},
			)
			require.NoError(t, errNormalization)

			fmt.Println(normalized)
		},
	)

	t.Run(
		"2. nested CSS",
		func(t *testing.T) {
			normalized, errNormalization := BeautifyCSS(
				&ParamsUpdateCSS{
					CSS: `
	.parent {
     .child {
    property: value;
  }
}
	`,
					ParamsSpaces: ParamsSpaces{
						NumberSpaces:              5,
						IncrementWithNumberSpaces: 2,
					},
				},
			)

			require.NoError(t, errNormalization)
			fmt.Println(normalized)
		},
	)

	t.Run(
		"3. CSS with comments",
		func(t *testing.T) {
			normalized, errNormalization := BeautifyCSS(
				&ParamsUpdateCSS{
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
					ParamsSpaces: ParamsSpaces{
						NumberSpaces:              5,
						IncrementWithNumberSpaces: 2,
					},
				},
			)

			require.NoError(t, errNormalization)
			fmt.Println(normalized)
		},
	)
}

func TestMinifyCSS(t *testing.T) {
	t.Run(
		"1. simple CSS",
		func(t *testing.T) {
			normalized, errNormalization := MinifyCSS(
				&ParamsUpdateCSS{
					CSS: `
	h1 {
font-size: 75px;
line-height: 1.15;
}

h2 {
font-size: 45px;
}
	`,
					ParamsSpaces: ParamsSpaces{
						NumberSpaces:              5,
						IncrementWithNumberSpaces: 2,
					},
				},
			)
			require.NoError(t, errNormalization)

			fmt.Println(normalized)
		},
	)

	t.Run(
		"2. nested CSS",
		func(t *testing.T) {
			normalized, errNormalization := MinifyCSS(
				&ParamsUpdateCSS{
					CSS: `
	.parent {
     .child {
    property: value;
  }
}
	`,
					ParamsSpaces: ParamsSpaces{
						NumberSpaces:              5,
						IncrementWithNumberSpaces: 2,
					},
				},
			)

			require.NoError(t, errNormalization)
			fmt.Println(normalized)
		},
	)

	t.Run(
		"3. CSS with comments",
		func(t *testing.T) {
			normalized, errNormalization := MinifyCSS(
				&ParamsUpdateCSS{
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
					ParamsSpaces: ParamsSpaces{
						NumberSpaces:              5,
						IncrementWithNumberSpaces: 2,
					},
				},
			)

			require.NoError(t, errNormalization)
			fmt.Println(normalized)
		},
	)
}
