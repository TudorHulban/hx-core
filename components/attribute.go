package hxcomponents

import (
	"database/sql"
	"io"
	"text/template"
)

type attribute struct {
	name  string
	value sql.NullString
}

func (a attribute) isAttribute() bool {
	return true
}

func (a attribute) Render(w io.Writer) error {
	if !a.value.Valid {
		_, err := w.Write(
			[]byte(_blank + a.name),
		)

		return err
	}

	_, errWrite := w.Write(
		[]byte(_blank + a.name + `="` + template.HTMLEscapeString(a.value.String) + `"`),
	)

	return errWrite
}
