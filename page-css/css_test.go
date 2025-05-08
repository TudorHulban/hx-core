package pagecss

import (
	"fmt"
	"strings"
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
			"2. CSS common only with blank added for testing the CSS normalization",
			func() *CSSElement {
				return &CSSElement{
					CSSAllMedias: " body{margin:0;}",
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

				output := NormalizeCSS(
					page.GetCSSAccurate(),
				)

				require.Equal(t,
					tt.want,
					output,

					fmt.Sprintf(
						"GetCSS() = %v, want %v",

						output,
						tt.want,
					),
				)
			},
		)
	}
}

func TestTwoElementsCSSPage(t *testing.T) {
	type contains struct {
		elements []string
		inPart   int
	}

	tests := []struct {
		name     string
		input1   CSS
		input2   CSS
		want     string
		contains contains
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
			contains{},
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
			contains{},
		},
		{
			"3a. CSS mixt with same inflexion point - mobile first",
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
			contains{
				elements: []string{
					"margin:0",
					"padding:0",
				},
				inPart: 1,
			},
		},
		{
			"3b. CSS mixt with same inflexion point - desktop first",
			func() *CSSElement {
				return &CSSElement{
					DesktopFirst: true,

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
					DesktopFirst: true,

					CSSResponsive: []CSSMedia{
						{
							InflexionPointPX: 768,
							CSS:              "body{padding:0;}",
						},
					},
				}
			},
			"@media (min-width: 768px) {\nbody{padding:0;}\nbody{margin:0;}\n}",
			contains{
				elements: []string{
					"margin:0",
					"padding:0",
				},
				inPart: 1,
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
			contains{},
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
			"body{color:white;}\n@media (min-width: 768px) and (max-width: 1365px) {\nbody{padding:0;}\n}\n@media (min-width: 1366px) {\nbody{margin: 5;}\n}",
			contains{},
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
			contains{
				elements: []string{
					"margin:0",
					"padding:0",
				},
				inPart: 1,
			},
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

				if len(tt.contains.elements) == 0 {
					fmt.Printf(
						"want:\n%s\n",
						tt.want,
					)

					t.FailNow()
				}

				outputParts := strings.Split(output, "@media")

				hxhelpers.ForEachTest(
					t,
					tt.contains.elements,
					func(value string, t *testing.T) {
						require.Contains(t,
							outputParts[tt.contains.inPart],
							value,

							fmt.Sprintf(
								"could not find '%s' in '%s'",

								value,
								outputParts[tt.contains.inPart],
							),
						)
					},
				)
			},
		)
	}
}
