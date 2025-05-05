package pagecss

import (
	"io"
	"slices"
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

// func (page CSSPage) GetCSS() string {
// 	cssCommon := make([]string, 0)
// 	cssResponsiveMap := make(map[uint16][]string)
// 	inflexionPoints := make([]uint16, 0)

// 	for cssFunc := range page {
// 		element := (*cssFunc)()

// 		if element.CSSAllMedias != "" {
// 			cssCommon = append(cssCommon, element.CSSAllMedias)
// 		}

// 		for _, media := range element.CSSResponsive {
// 			cssResponsiveMap[media.InflexionPointPX] = append(
// 				cssResponsiveMap[media.InflexionPointPX],
// 				media.CSS,
// 			)

// 			if !slices.Contains(inflexionPoints, media.InflexionPointPX) {
// 				inflexionPoints = append(
// 					inflexionPoints,
// 					media.InflexionPointPX,
// 				)
// 			}
// 		}
// 	}

// 	slices.Sort(inflexionPoints)

// 	cssResponsive := make([]string, 0)

// 	for ix, point := range inflexionPoints {
// 		cssRules := cssResponsiveMap[point]
// 		sort.Strings(cssRules)
// 		css := strings.Join(cssRules, "\n")
// 		mediaQuery := "@media (min-width: " + strconv.Itoa(int(point)) + "px)"

// 		if ix < len(inflexionPoints)-1 {
// 			nextPoint := inflexionPoints[ix+1]
// 			mediaQuery = mediaQuery + " and (max-width: " + strconv.Itoa(int(nextPoint)-1) + "px)"
// 		}

// 		cssResponsive = append(
// 			cssResponsive,
// 			mediaQuery+" {\n"+css+"\n}",
// 		)
// 	}

// 	allCSS := strings.Join(cssCommon, "\n")

// 	if len(cssResponsive) > 0 {
// 		if allCSS != "" {
// 			allCSS = allCSS + "\n"
// 		}

// 		allCSS = allCSS + strings.Join(cssResponsive, "\n")
// 	}

// 	return allCSS
// }

func (page CSSPage) GetCSS() string {
	cssCommon := make([]string, 0, len(page))
	cssResponsiveMap := make(map[uint16]map[string]map[string]string)
	inflexionPoints := make([]uint16, 0, len(page))

	cssFuncs := make([]CSS, 0, len(page))
	for cssFunc := range page {
		f := (*cssFunc)

		cssFuncs = append(cssFuncs, f)
	}

	for _, cssFunc := range cssFuncs {
		element := cssFunc()
		if element == nil {
			continue
		}

		if element.CSSAllMedias != "" {
			cssCommon = append(
				cssCommon,
				element.CSSAllMedias,
			)
		}

		for _, media := range element.CSSResponsive {
			if media.CSS == "" {
				continue
			}

			if _, exists := cssResponsiveMap[media.InflexionPointPX]; !exists {
				cssResponsiveMap[media.InflexionPointPX] = make(map[string]map[string]string)

				inflexionPoints = append(
					inflexionPoints,
					media.InflexionPointPX,
				)
			}

			parts := strings.SplitN(media.CSS, "{", 2)
			if len(parts) != 2 {
				continue
			}

			selector := strings.TrimSpace(parts[0])
			properties := strings.TrimRight(strings.TrimSpace(parts[1]), "}")
			propPairs := strings.Split(properties, ";")

			if _, exists := cssResponsiveMap[media.InflexionPointPX][selector]; !exists {
				cssResponsiveMap[media.InflexionPointPX][selector] = make(map[string]string)
			}

			for _, pair := range propPairs {
				pair = strings.TrimSpace(pair)
				if pair == "" {
					continue
				}

				propParts := strings.SplitN(pair, ":", 2)
				if len(propParts) != 2 {
					continue
				}

				property := strings.TrimSpace(propParts[0])
				cssResponsiveMap[media.InflexionPointPX][selector][property] = media.CSS
			}
		}
	}

	slices.Sort(inflexionPoints)

	cssResponsive := make([]string, 0, len(inflexionPoints))
	var builder strings.Builder

	for ix, point := range inflexionPoints {
		// Collect rules for this media query: prioritize current point, then higher points
		ruleMap := make(map[string]map[string]string)

		// First, add rules from higher points (in reverse order to prioritize later points)
		for i := len(inflexionPoints) - 1; i > ix; i-- {
			p := inflexionPoints[i]

			for selector, props := range cssResponsiveMap[p] {
				if _, exists := ruleMap[selector]; !exists {
					ruleMap[selector] = make(map[string]string)
				}

				for prop, rule := range props {
					ruleMap[selector][prop] = rule
				}
			}
		}

		// Then, add rules from current point (overrides higher points)
		for selector, props := range cssResponsiveMap[point] {
			if _, exists := ruleMap[selector]; !exists {
				ruleMap[selector] = make(map[string]string)
			}

			for prop, rule := range props {
				ruleMap[selector][prop] = rule
			}
		}

		// Convert to ordered rules, prioritizing later functions
		rules := make([]string, 0)
		seen := make(map[string]bool)

		for i := len(cssFuncs) - 1; i >= 0; i-- { // Reverse to prioritize later functions
			cssFunc := cssFuncs[i]

			element := cssFunc()
			if element == nil {
				continue
			}

			for _, media := range element.CSSResponsive {
				if media.InflexionPointPX < point {
					continue
				}

				parts := strings.SplitN(media.CSS, "{", 2)
				if len(parts) != 2 {
					continue
				}

				selector := strings.TrimSpace(parts[0])
				properties := strings.TrimRight(strings.TrimSpace(parts[1]), "}")
				propPairs := strings.Split(properties, ";")

				for _, pair := range propPairs {
					pair = strings.TrimSpace(pair)
					if pair == "" {
						continue
					}

					propParts := strings.SplitN(pair, ":", 2)
					if len(propParts) != 2 {
						continue
					}

					property := strings.TrimSpace(propParts[0])
					if ruleMap[selector][property] == media.CSS && !seen[media.CSS] {
						seen[media.CSS] = true
						rules = append(rules, media.CSS)
					}
				}
			}
		}

		css := strings.Join(rules, "\n")
		if css == "" {
			continue
		}

		builder.Reset()
		builder.WriteString("@media (min-width: ")
		builder.WriteString(strconv.Itoa(int(point)))
		builder.WriteString("px)")

		if ix < len(inflexionPoints)-1 {
			nextPoint := inflexionPoints[ix+1]

			builder.WriteString(" and (max-width: ")
			builder.WriteString(strconv.Itoa(int(nextPoint) - 1))
			builder.WriteString("px)")
		}

		builder.WriteString(" {\n")
		builder.WriteString(css)
		builder.WriteString("\n}")

		cssResponsive = append(cssResponsive, builder.String())
	}

	builder.Reset()

	for i, css := range cssCommon {
		builder.WriteString(css)

		if i < len(cssCommon)-1 || len(cssResponsive) > 0 {
			builder.WriteString("\n")
		}
	}

	if len(cssResponsive) > 0 {
		builder.WriteString(strings.Join(cssResponsive, "\n"))
	}

	return builder.String()
}

func (page CSSPage) GetCSSTo(w io.Writer) (int, error) {
	return w.Write(
		[]byte(page.GetCSS()),
	)
}
