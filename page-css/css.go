package pagecss

import (
	"io"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type CSSMedia struct {
	CSS              string
	InflexionPointPX uint16
}

type CSSElement struct {
	CSSAllMedias  string
	CSSResponsive []CSSMedia
}

type CSS func() *CSSElement

type CSSPage map[*CSS]bool

func NewCSSPage(css ...func() *CSSElement) CSSPage {
	var result CSSPage = map[*CSS]bool{}

	for _, f := range css {
		var buf CSS = f

		if _, exists := result[&buf]; exists {
			continue
		}

		result[&buf] = true
	}

	return result
}

func (page CSSPage) GetCSS() string {
	cssCommon := make([]string, 0)
	cssResponsiveMap := make(map[uint16][]string)
	inflexionPoints := make([]uint16, 0)

	for cssFunc := range page {
		element := (*cssFunc)()

		if element.CSSAllMedias != "" {
			cssCommon = append(cssCommon, element.CSSAllMedias)
		}

		for _, media := range element.CSSResponsive {
			cssResponsiveMap[media.InflexionPointPX] = append(
				cssResponsiveMap[media.InflexionPointPX],
				media.CSS,
			)

			if !slices.Contains(inflexionPoints, media.InflexionPointPX) {
				inflexionPoints = append(
					inflexionPoints,
					media.InflexionPointPX,
				)
			}
		}
	}

	slices.Sort(inflexionPoints)

	cssResponsive := make([]string, 0)

	for ix, point := range inflexionPoints {
		cssRules := cssResponsiveMap[point]
		sort.Strings(cssRules)
		css := strings.Join(cssRules, "\n")
		mediaQuery := "@media (min-width: " + strconv.Itoa(int(point)) + "px)"

		if ix < len(inflexionPoints)-1 {
			nextPoint := inflexionPoints[ix+1]
			mediaQuery = mediaQuery + " and (max-width: " + strconv.Itoa(int(nextPoint)-1) + "px)"
		}

		cssResponsive = append(
			cssResponsive,
			mediaQuery+" {\n"+css+"\n}",
		)
	}

	allCSS := strings.Join(cssCommon, "\n")

	if len(cssResponsive) > 0 {
		if allCSS != "" {
			allCSS = allCSS + "\n"
		}

		allCSS = allCSS + strings.Join(cssResponsive, "\n")
	}

	return allCSS
}

func (page CSSPage) GetCSSTo(w io.Writer) (int, error) {
	return w.Write(
		[]byte(page.GetCSS()),
	)
}
