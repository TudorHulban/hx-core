package pagecss

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

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

			t.Log(normalized)

			expected := `
h1 {
  font-size:  75px;
  line-height:  1.15;
}

h2 {
  font-size:  45px;
}
			`

			require.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(normalized))
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

			t.Log(normalized)

			expected := `
.parent {
  .child{
    property:  value;
  }
}
`

			require.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(normalized))

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

			t.Log(normalized)

			expected := `
/* Input CSS */
.class {
  /* This comment contains a closing brace } that confuses the logic */
  property:  value;
}

/* Output using the provided logic */
.class{
  /* This comment contains a closing brace } that confuses the logic */
  property:  value;
  /* This line might not be indented */
}
			`

			require.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(normalized))
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

			t.Log(normalized)

			expected := `h1{font-size:75px;line-height:1.15}h2{font-size:45px}`

			require.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(normalized))
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

			t.Log(normalized)

			expected := `.parent{.child{property:value}}`

			require.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(normalized))
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

			t.Log(normalized)

			expected := `.class{property:value}.class{property:value}`

			require.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(normalized))
		},
	)
}
