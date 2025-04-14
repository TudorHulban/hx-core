package hxprimitives

import "database/sql"

func Attr(name string) attribute {
	return attribute{
		name: name,
	}
}

func AttrIf(condition bool, name string) attribute {
	if condition {
		return attribute{
			name: name,
		}
	}

	return attribute{}
}

func AttrWithName(name string) attribute {
	return AttrWithValue(
		"name",
		name,
	)
}

func AttrWithValue(name, value string) attribute {
	return attribute{
		name: name,
		value: sql.NullString{
			Valid:  true,
			String: value,
		},
	}
}

func AttrWithValueIf(condition bool, name, value string) attribute {
	if condition {
		return attribute{
			name: name,
			value: sql.NullString{
				Valid:  true,
				String: value,
			},
		}
	}

	return attribute{}
}

func AttrCSS(css string) attribute {
	return AttrWithValue(
		"style",
		css,
	)
}

func AttrClass(class string) attribute {
	return AttrWithValue(
		"class",
		class,
	)
}

func AttrClassLength(class string) attribute {
	if len(class) > 0 {
		return AttrWithValue(
			"class",
			class,
		)
	}

	return attribute{}
}

func AttrID(value string) attribute {
	return AttrWithValue(
		"id",
		value,
	)
}

func AttrIDIf(condition bool, value string) attribute {
	if condition {
		return AttrWithValue(
			"id",
			value,
		)
	}

	return attribute{}
}

// AttrIDLength returns css id attribute if passed value valid.
func AttrIDLength(value string) attribute {
	if len(value) > 0 {
		return AttrWithValue(
			"id",
			value,
		)
	}

	return attribute{}
}

func AttrType(class string) attribute {
	return AttrWithValue(
		"type",
		class,
	)
}

func LabelFor(text string) attribute {
	return AttrWithValue(
		"for",
		text,
	)
}

func Lang(language string) attribute {
	return AttrWithValue(
		"lang",
		language,
	)
}

func Charset(charset string) attribute {
	return AttrWithValue(
		"charset",
		charset,
	)
}

func Content(content string) attribute {
	return AttrWithValue(
		"content",
		content,
	)
}

func Href(ref string) attribute {
	return AttrWithValue(
		"href",
		ref,
	)
}

func MethodPOST() attribute {
	return AttrWithValue(
		"method",
		"post",
	)
}

func Name(name string) attribute {
	return AttrWithValue(
		"name",
		name,
	)
}

func OnClick(toRun string) attribute {
	return AttrWithValue(
		"onclick",
		toRun,
	)
}

func Rel(relationship string) attribute {
	return AttrWithValue(
		"rel",
		relationship,
	)
}

func Src(source string) attribute {
	return AttrWithValue(
		"src",
		source,
	)
}
