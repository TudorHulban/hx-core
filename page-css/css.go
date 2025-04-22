package pagecss

import (
	"slices"
	"strconv"
	"strings"
)

type CSSMedia struct {
	InflexionPointPX uint16
	CSS              string
}

type CSSElement struct {
	CSSAllMedias  string
	CSSResponsive []CSSMedia
}

type CSS func() CSSElement

type CSSPage map[*CSS]bool

func (page CSSPage) GetCSS() string {
	cssCommon := make([]string, 0)
	cssResponsiveMap := make(map[uint16][]string)
	inflexionPoints := []uint16{}

	for cssFunc := range page {
		element := (*cssFunc)()

		if element.CSSAllMedias != "" {
			cssCommon = append(
				cssCommon,
				element.CSSAllMedias,
			)
		}

		for _, media := range element.CSSResponsive {
			cssResponsiveMap[media.InflexionPointPX] = append(
				cssResponsiveMap[media.InflexionPointPX],
				media.CSS,
			)

			var found bool

			if slices.Contains(inflexionPoints, media.InflexionPointPX) {
				found = true
			}

			if !found {
				inflexionPoints = append(inflexionPoints, media.InflexionPointPX)
			}
		}
	}

	slices.Sort(inflexionPoints)

	cssResponsive := []string{}

	for _, point := range inflexionPoints {
		css := strings.Join(cssResponsiveMap[point], "\n")

		cssResponsive = append(
			cssResponsive,
			"@media (min-width: "+strconv.Itoa(int(point))+"px) {\n"+css+"\n}",
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
