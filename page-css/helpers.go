package pagecss

import (
	"bytes"
	"errors"
	"io"
	"strings"

	"github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/css"
)

type ParamsNormalizeCSS struct {
	CSS                   string
	CSSNumberSpaces       uint8
	IncrementNumberSpaces uint8
}

func NormalizeCSS(params *ParamsNormalizeCSS) (string, error) {
	if params == nil {
		return "",
			errors.New("params cannot be nil")
	}

	// Pre-calculate indentation strings
	baseIndent := strings.Repeat(" ", int(params.CSSNumberSpaces))
	nestedIndentUnit := strings.Repeat(" ", int(params.IncrementNumberSpaces))

	// Create a lexer to analyze the CSS
	input := parse.NewInput(bytes.NewReader([]byte(params.CSS)))
	lexer := css.NewLexer(input)

	var buf bytes.Buffer
	buf.Grow(len(params.CSS) * 2) // Allocate enough space

	indentLevel := 0

	var inRuleset, startOfLine, pendingSelector, seenColon bool

	for {
		tokenType, text := lexer.Next()
		if tokenType == css.ErrorToken {
			if lexer.Err() != io.EOF {
				return "",
					lexer.Err()
			}

			break
		}

		currentIndent := strings.Repeat(nestedIndentUnit, indentLevel)

		switch tokenType {
		case css.IdentToken:
			if startOfLine && inRuleset && !seenColon {
				// This is likely a nested selector
				buf.WriteString(currentIndent)
				buf.Write(text)
				buf.WriteByte(' ')

				pendingSelector = true
			} else if inRuleset && !seenColon {
				// This is a property name
				if !pendingSelector {
					buf.WriteString(currentIndent)
					buf.WriteString(baseIndent) // Apply base indentation for declarations
				}

				buf.Write(text)

				pendingSelector = false
			} else if seenColon {
				// This is likely a property value
				buf.Write(text)
				buf.WriteByte(' ')
			} else {
				// This is a top-level selector
				if startOfLine {
					buf.WriteString(currentIndent)
				}

				buf.Write(text)
				buf.WriteByte(' ')

				pendingSelector = true
			}
			startOfLine = false

		case css.DelimToken:
			// Handle delimiters, including the dot for class selectors
			if string(text) == "." {
				if startOfLine && inRuleset {
					buf.WriteString(currentIndent)
					startOfLine = false
				}

				buf.WriteByte('.')

				pendingSelector = true
			} else {
				buf.Write(text)
			}

		case css.ColonToken:
			seenColon = true

			buf.WriteString(": ")

			startOfLine = false

		case css.SemicolonToken:
			buf.WriteByte(';')
			buf.WriteByte('\n')

			seenColon = false
			startOfLine = true
			pendingSelector = false

		case css.LeftBraceToken:
			inRuleset = true
			pendingSelector = false

			buf.WriteByte('{')
			buf.WriteByte('\n')
			indentLevel++
			startOfLine = true

		case css.RightBraceToken:
			if indentLevel > 0 {
				indentLevel--
			}
			currentIndent = strings.Repeat(nestedIndentUnit, indentLevel)

			buf.WriteString(currentIndent)
			buf.WriteByte('}')
			buf.WriteByte('\n')

			startOfLine = true

			if indentLevel == 0 {
				inRuleset = false
			}

		case css.NumberToken, css.PercentageToken, css.DimensionToken:
			// Handle numbers and dimensions (like px, em, etc.)
			buf.Write(text)
			startOfLine = false

		case css.CommaToken:
			buf.WriteByte(',')
			buf.WriteByte(' ')

			startOfLine = false

		case css.WhitespaceToken:
			// Only add a space if not at start of line and not after certain tokens
			if !startOfLine && !pendingSelector {
				buf.WriteByte(' ')
			}

		case css.HashToken:
			// Handle ID selectors
			if startOfLine {
				buf.WriteString(currentIndent)
			}
			buf.Write(text)
			buf.WriteByte(' ')
			startOfLine = false
			pendingSelector = true

		case css.AtKeywordToken:
			// Handle at-rules like @media
			buf.WriteString(currentIndent)
			buf.Write(text)
			buf.WriteByte(' ')
			startOfLine = false

		case css.CommentToken:
			buf.WriteString(currentIndent)
			buf.Write(text)
			buf.WriteByte('\n')
			startOfLine = true

		default:
			// Handle other tokens
			buf.Write(text)
			startOfLine = false
		}
	}

	return buf.String(), nil
}
