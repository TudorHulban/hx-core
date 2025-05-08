package pagecss

import (
	"fmt"
	"testing"

	hxhelpers "github.com/TudorHulban/hx-core/helpers"
	"github.com/stretchr/testify/require"
)

func TestOneElementCSSPage(t *testing.T) {
	tests := []struct {
		name  string
		input CSS
		want  string
	}{
		{
			"1. CSS empty",
			func() *CSSElement {
				return &CSSElement{}
			},
			"",
		},
		{
			"2. CSS common only",
			func() *CSSElement {
				return &CSSElement{
					CSSAllMedias: "body{margin:0;}",
				}
			},
			"body{margin:0;}",
		},
		{
			"3. CSS responsive only",
			func() *CSSElement {
				return &CSSElement{
					CSSResponsive: []CSSMedia{
						{
							InflexionPointPX: 768,
							CSS:              "body{margin:0;}",
						},
					},
				}
			},
			"@media (min-width: 768px) {\nbody{margin:0;}\n}",
		},
		{
			"4. CSS common and responsive",
			func() *CSSElement {
				return &CSSElement{
					CSSAllMedias: "body{}",
					CSSResponsive: []CSSMedia{
						{
							InflexionPointPX: 768,
							CSS:              "body{margin: 0;}",
						},
					},
				}
			},
			"body{}\n@media (min-width: 768px) {\nbody{margin: 0;}\n}",
		},
		{
			"5. Multiple inflexion points",
			func() *CSSElement {
				return &CSSElement{
					CSSAllMedias: "body{}",
					CSSResponsive: []CSSMedia{
						{
							InflexionPointPX: 768,
							CSS:              "body{margin: 0;}",
						},
						{
							InflexionPointPX: 1366,
							CSS:              "body{margin: 5;}",
						},
						{
							InflexionPointPX: 768,
							CSS:              "body{padding: 0;}",
						},
					},
				}
			},
			"body{}\n@media (min-width: 768px) and (max-width: 1365px) {\nbody{margin: 0;}\nbody{padding: 0;}\n}\n@media (min-width: 1366px) {\nbody{margin: 5;}\n}",
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				var page CSSPage = map[*CSS]bool{
					&tt.input: true,
				}

				require.Equal(t,
					tt.want,
					page.GetCSSAccurate(),

					fmt.Sprintf(
						"GetCSS() = %v, want %v",

						page.GetCSSAccurate(),
						tt.want,
					),
				)
			},
		)
	}
}

func TestTwoElementsCSSPage(t *testing.T) {
	tests := []struct {
		name        string
		input1      CSS
		input2      CSS
		want        string
		butContains []string
	}{
		{
			"1. One CSS empty",
			func() *CSSElement {
				return &CSSElement{}
			},
			func() *CSSElement {
				return &CSSElement{
					CSSAllMedias: "body{}",
				}
			},
			"body{}",
			nil,
		},
		{
			"2. CSS common and responsive",
			func() *CSSElement {
				return &CSSElement{
					CSSAllMedias: "body{margin:0;}",
				}
			},
			func() *CSSElement {
				return &CSSElement{
					CSSResponsive: []CSSMedia{
						{
							InflexionPointPX: 768,
							CSS:              "body{padding:0;}",
						},
					},
				}
			},
			"body{margin:0;}\n@media (min-width: 768px) {\nbody{padding:0;}\n}",
			nil,
		},
		{
			"3. CSS mixt with same inflexion point",
			func() *CSSElement {
				return &CSSElement{
					CSSResponsive: []CSSMedia{
						{
							InflexionPointPX: 768,
							CSS:              "body{margin:0;}",
						},
					},
				}
			},
			func() *CSSElement {
				return &CSSElement{
					CSSResponsive: []CSSMedia{
						{
							InflexionPointPX: 768,
							CSS:              "body{padding:0;}",
						},
					},
				}
			},
			"@media (min-width: 768px) {\nbody{padding:0;}\nbody{margin:0;}\n}",
			[]string{
				"margin:0",
				"padding:0",
			},
		},
		{
			"4. CSS mixt with one lower inflexion point",
			func() *CSSElement {
				return &CSSElement{
					CSSAllMedias: "body{color:white;}",
					CSSResponsive: []CSSMedia{
						{
							InflexionPointPX: 768,
							CSS:              "body{margin:0;}",
						},
					},
				}
			},
			func() *CSSElement {
				return &CSSElement{
					CSSResponsive: []CSSMedia{
						{
							InflexionPointPX: 768,
							CSS:              "body{padding:0;}",
						},
						{
							InflexionPointPX: 1366,
							CSS:              "body{margin: 5;}",
						},
					},
				}
			},
			"body{color:white;}\n@media (min-width: 768px) and (max-width: 1365px) {\nbody{padding:0;}\nbody{margin:0;}\n}\n@media (min-width: 1366px) {\nbody{margin: 5;}\n}",
			[]string{
				"margin:0",
				"padding:0",
			},
		},
		{
			"5. CSS mixt with one higher inflexion point",
			func() *CSSElement {
				return &CSSElement{
					CSSAllMedias: "body{color:white;}",
					CSSResponsive: []CSSMedia{
						{
							InflexionPointPX: 1366,
							CSS:              "body{margin:0;}",
						},
					},
				}
			},
			func() *CSSElement {
				return &CSSElement{
					CSSResponsive: []CSSMedia{
						{
							InflexionPointPX: 768,
							CSS:              "body{padding:0;}",
						},
						{
							InflexionPointPX: 1366,
							CSS:              "body{margin: 5;}",
						},
					},
				}
			},
			"body{color:white;}\n@media (min-width: 768px) and (max-width: 1365px) {\nbody{padding:0;}\n}\n@media (min-width: 1366px) {\nbody{margin:0;}\nbody{margin: 5;}\n}",
			nil,
		},
		{
			"6. CSS mixt many inflexion points",
			func() *CSSElement {
				return &CSSElement{
					CSSAllMedias: "body{color:white;}",
					CSSResponsive: []CSSMedia{
						{
							InflexionPointPX: 1900,
							CSS:              "body{margin: 15;}",
						},
						{
							InflexionPointPX: 768,
							CSS:              "body{margin:0;}",
						},
					},
				}
			},
			func() *CSSElement {
				return &CSSElement{
					CSSResponsive: []CSSMedia{
						{
							InflexionPointPX: 768,
							CSS:              "body{padding:0;}",
						},
						{
							InflexionPointPX: 1366,
							CSS:              "body{margin: 5;}",
						},
					},
				}
			},
			"body{color:white;}\n@media (min-width: 768px) and (max-width: 1365px) {\nbody{margin:0;}\nbody{padding:0;}\n}\n@media (min-width: 1366px) and (max-width: 1899px) {\nbody{margin: 5;}\n}\n@media (min-width: 1900px) {\nbody{margin: 15;}\n}",
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				var page CSSPage = map[*CSS]bool{
					&tt.input1: true,
					&tt.input2: true,
				}

				output := page.GetCSSAccurate()

				fmt.Println(tt.name)
				fmt.Printf(
					"output:\n%s\n",
					output,
				)

				if tt.want == output {
					return
				}

				if len(tt.butContains) == 0 {
					fmt.Printf(
						"want:\n%s\n",
						tt.want,
					)

					t.FailNow()
				}

				hxhelpers.ForEachTest(
					t,
					tt.butContains,
					func(value string, t *testing.T) {
						require.Contains(t,
							output,
							value,
						)
					},
				)
			},
		)
	}
}
