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

	t.Run(
		"2. media query",
		func(t *testing.T) {
			normalized, errNormalization := BeautifyCSS(
				&ParamsUpdateCSS{
					CSS: `
@media (min-width :768px){
h1  {
font-size :  75px;
    line-height :  1.15;
  }
	
   h2  {
    font-size :  45px;
  }
  h3  {
    font-size :  33px;
    line-height :  1.25;
  }
  table  th,         table        td {
    padding :  18px 20px;
  }
  #form-footer-info  {
    flex-direction :  column ;
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
}
