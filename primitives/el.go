package hxprimitives

import (
	"database/sql"
	"io"
)

func El(name string, nodes ...Node) Writer {
	return func(w io.Writer) error {
		return renderNodes(
			w,

			sql.NullString{
				Valid:  true,
				String: name,
			},
			nodes...,
		)
	}
}

func ElWId(name, cssID string, nodes ...Node) Writer {
	return func(w io.Writer) error {
		return renderNodesWithCSSId(
			w,

			name,
			cssID,
			nodes...,
		)
	}
}
