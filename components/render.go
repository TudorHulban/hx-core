package hxcomponents

import (
	"database/sql"
	"io"

	goerrors "github.com/TudorHulban/go-errors"
)

func renderNodes(w io.Writer, elName sql.NullString, nodes ...Node) error {
	if w == nil {
		return goerrors.ErrNilInput{
			InputName: "writer",
		}
	}

	if len(nodes) == 0 && !isTagless(elName.String) {
		return goerrors.ErrNilInput{
			InputName: "nodes",
		}
	}

	if elName.Valid {
		_, _ = w.Write([]byte("<" + elName.String))

		for _, node := range nodes {
			if node == nil {
				continue
			}

			if node.isAttribute() {
				if errRender := node.Render(w); errRender != nil {
					return errRender
				}
			}
		}

		_, _ = w.Write([]byte(">"))

		if isTagless(elName.String) {
			return nil
		}
	}

	for _, childNode := range nodes {
		if childNode == nil || childNode.isAttribute() {
			continue
		}

		if errRender := childNode.Render(w); errRender != nil {
			return errRender
		}
	}

	if elName.Valid && !isTagless(elName.String) {
		_, _ = w.Write([]byte("</" + elName.String + ">"))
	}

	return nil
}

func renderNodesWithCSSId(w io.Writer, elName, cssID string, nodes ...Node) error {
	if w == nil {
		return goerrors.ErrNilInput{
			InputName: "writer",
		}
	}

	if len(nodes) == 0 && !isTagless(elName) {
		return goerrors.ErrNilInput{
			InputName: "nodes",
		}
	}

	_, _ = w.Write([]byte("<" + elName))

	_ = AttrIDLength(cssID).Render(w)

	for _, node := range nodes {
		if node == nil {
			continue
		}

		if node.isAttribute() {
			if errRender := node.Render(w); errRender != nil {
				return errRender
			}
		}
	}

	_, _ = w.Write([]byte(">"))

	if isTagless(elName) {
		return nil
	}

	for _, node := range nodes {
		if node == nil || node.isAttribute() {
			continue
		}

		if errRender := node.Render(w); errRender != nil {
			return errRender
		}
	}

	if !isTagless(elName) {
		_, _ = w.Write([]byte("</" + elName + ">"))
	}

	return nil
}
