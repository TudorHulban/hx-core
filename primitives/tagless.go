package hxprimitives

var tagless = map[string]struct{}{
	"area":   {},
	"br":     {},
	"hr":     {},
	"img":    {},
	"input":  {},
	"link":   {},
	"meta":   {},
	"source": {},
}

func isTagless(elName string) bool {
	_, exists := tagless[elName]

	return exists
}
