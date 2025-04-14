package hxhelpers

import (
	"strings"
)

func SanitizeCSSId(element string) string {
	return "#" + strings.TrimLeft(element, "#")
}

func SanitizeCSSIds(elements []string) string {
	return strings.Join(
		ForEachValue(elements, SanitizeCSSId),
		",",
	)
}

func SanitizeCSSIdsSlice(elements ...string) []string {
	return ForEachValue(
		elements,
		SanitizeCSSId,
	)
}
