package pagecss

import (
	"bytes"
	"errors"
	"io"
	"strings"

	"github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/css"
)

type ParamsSpaces struct {
	NumberSpaces              uint8
	IncrementWithNumberSpaces uint8
}

type ParamsUpdateCSS struct {
	CSS string

	ParamsSpaces
}

type paramsNormalizeCSS struct {
	ParamsUpdateCSS

	isBeautifier bool
}

func normalizeCSS(params *paramsNormalizeCSS) (string, error) {
	if params == nil {
		return "",
			errors.New("params cannot be nil")
	}

	lexer := css.NewLexer(
		parse.NewInput(bytes.NewReader([]byte(params.CSS))),
	)

	var buf bytes.Buffer
	buf.Grow(len(params.CSS) * 2) // Allocate enough space

	indentLevel := 0
	var inRuleset, startOfLine, pendingSelector, seenColon, lastWasRightBrace bool
	var inSelector bool = true // Start true as we typically begin with selectors

	// Only use indentation if we are in beautifier mode
	var baseIndent, nestedIndentUnit string

	if params.isBeautifier {
		baseIndent = strings.Repeat(
			" ",
			int(params.ParamsSpaces.NumberSpaces),
		)
		nestedIndentUnit = strings.Repeat(
			" ",
			int(params.ParamsSpaces.IncrementWithNumberSpaces),
		)

		startOfLine = true // Initialize only for beautifier
	}

	for {
		tokenType, text := lexer.Next()
		if tokenType == css.ErrorToken {
			if lexer.Err() != io.EOF {
				return "", lexer.Err()
			}

			break
		}

		var currentIndent string

		if params.isBeautifier {
			currentIndent = strings.Repeat(nestedIndentUnit, indentLevel)
		}

		switch tokenType {
		case css.IdentToken:
			if params.isBeautifier {
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
				lastWasRightBrace = false
			} else {
				// Minification mode
				buf.Write(text)
			}

		case css.DelimToken:
			delimText := string(text)
			if delimText == "." {
				if params.isBeautifier && startOfLine && inRuleset {
					buf.WriteString(currentIndent)
					startOfLine = false
				}
				buf.WriteByte('.')

				pendingSelector = true
			} else if delimText == ":" {
				// Just write the colon without any space manipulation
				buf.WriteByte(':')

				// Only set seenColon for property:value pairs, not for selectors
				if !inSelector && inRuleset {
					seenColon = true
					if params.isBeautifier {
						buf.WriteByte(' ') // Space after property colon (not in selectors)
					}
				}
			} else {
				buf.Write(text)
			}

			lastWasRightBrace = false

		case css.ColonToken:
			seenColon = true
			if params.isBeautifier && !inSelector {
				buf.WriteString(": ")
				startOfLine = false
			} else {
				buf.WriteByte(':')
			}

			lastWasRightBrace = false

		case css.SemicolonToken:
			if params.isBeautifier {
				buf.WriteByte(';')
				buf.WriteByte('\n')
				seenColon = false
				startOfLine = true
				pendingSelector = false
			} else {
				// For minification, add semicolon
				buf.WriteByte(';')
				seenColon = false
			}

			lastWasRightBrace = false

		case css.LeftBraceToken:
			inRuleset = true
			pendingSelector = false
			inSelector = false // Inside braces, not in selector

			if params.isBeautifier {
				buf.WriteByte('{')
				buf.WriteByte('\n')
				indentLevel++
				startOfLine = true
			} else {
				buf.WriteByte('{')
			}

			lastWasRightBrace = false

		case css.RightBraceToken:
			if params.isBeautifier {
				if indentLevel > 0 {
					indentLevel--
				}
				currentIndent = strings.Repeat(nestedIndentUnit, indentLevel)
				buf.WriteString(currentIndent)
				buf.WriteByte('}')
				buf.WriteByte('\n')

				// Add an extra blank line after each CSS block (but not inside nested blocks)
				if indentLevel == 0 {
					buf.WriteByte('\n')
				}

				startOfLine = true
				lastWasRightBrace = true
			} else {
				// For minification: remove the last semicolon if present
				// Get current buffer content
				bufContent := buf.Bytes()

				// Check if the last character is a semicolon
				if len(bufContent) > 0 && bufContent[len(bufContent)-1] == ';' {
					// Remove only the semicolon
					buf.Truncate(buf.Len() - 1)
				}

				buf.WriteByte('}')
			}

			if indentLevel == 0 {
				inRuleset = false
				inSelector = true // Back to selector context outside of braces
			}

		case css.NumberToken, css.PercentageToken, css.DimensionToken:
			// Handle numbers and dimensions (like px, em, etc.)
			buf.Write(text)
			if params.isBeautifier {
				startOfLine = false
			}

			lastWasRightBrace = false

		case css.CommaToken:
			if params.isBeautifier {
				buf.WriteByte(',')
				buf.WriteByte(' ')
				startOfLine = false
			} else {
				buf.WriteByte(',')
			}

			lastWasRightBrace = false

		case css.WhitespaceToken:
			if params.isBeautifier {
				// Skip whitespace in selector context to avoid breaking pseudo-elements
				if !startOfLine && !inSelector {
					buf.WriteByte(' ')
				}
			}
			// In minification mode, skip all whitespace

		case css.HashToken:
			// Handle ID selectors
			if params.isBeautifier {
				if startOfLine {
					buf.WriteString(currentIndent)
				}
				buf.Write(text)
				buf.WriteByte(' ')

				startOfLine = false
				pendingSelector = true
			} else {
				buf.Write(text)
			}

			lastWasRightBrace = false

		case css.AtKeywordToken:
			// Handle at-rules like @media
			if params.isBeautifier {
				if lastWasRightBrace {
					// Add an extra newline before at-rules if we just closed a block
					buf.WriteByte('\n')
				}
				buf.WriteString(currentIndent)
				buf.Write(text)
				buf.WriteByte(' ')

				startOfLine = false
			} else {
				buf.Write(text)
			}

			lastWasRightBrace = false

		case css.CommentToken:
			if params.isBeautifier {
				buf.WriteString(currentIndent)
				buf.Write(text)
				buf.WriteByte('\n')

				startOfLine = true
			}
			// Skip comments in minification mode
			lastWasRightBrace = false

		default:
			// Handle other tokens
			buf.Write(text)
			if params.isBeautifier {
				startOfLine = false
			}
			lastWasRightBrace = false
		}
	}

	return buf.String(), nil
}

// func normalizeCSS(params *paramsNormalizeCSS) (string, error) {
// 	if params == nil {
// 		return "", errors.New("params cannot be nil")
// 	}

// 	// Create a lexer to analyze the CSS
// 	input := parse.NewInput(bytes.NewReader([]byte(params.CSS)))
// 	lexer := css.NewLexer(input)

// 	var buf bytes.Buffer
// 	buf.Grow(len(params.CSS) * 2) // Allocate enough space

// 	indentLevel := 0
// 	var inRuleset, startOfLine, pendingSelector, seenColon bool

// 	// Only use indentation if we're in beautifier mode
// 	var baseIndent, nestedIndentUnit string

// 	if params.isBeautifier {
// 		baseIndent = strings.Repeat(
// 			" ",
// 			int(params.ParamsSpaces.NumberSpaces),
// 		)
// 		nestedIndentUnit = strings.Repeat(
// 			" ",
// 			int(params.ParamsSpaces.IncrementWithNumberSpaces),
// 		)
// 		startOfLine = true // Initialize only for beautifier
// 	}

// 	for {
// 		tokenType, text := lexer.Next()
// 		if tokenType == css.ErrorToken {
// 			if lexer.Err() != io.EOF {
// 				return "", lexer.Err()
// 			}

// 			break
// 		}

// 		var currentIndent string

// 		if params.isBeautifier {
// 			currentIndent = strings.Repeat(nestedIndentUnit, indentLevel)
// 		}

// 		switch tokenType {
// 		case css.IdentToken:
// 			if params.isBeautifier {
// 				if startOfLine && inRuleset && !seenColon {
// 					// This is likely a nested selector
// 					buf.WriteString(currentIndent)
// 					buf.Write(text)
// 					buf.WriteByte(' ')
// 					pendingSelector = true
// 				} else if inRuleset && !seenColon {
// 					// This is a property name
// 					if !pendingSelector {
// 						buf.WriteString(currentIndent)
// 						buf.WriteString(baseIndent) // Apply base indentation for declarations
// 					}
// 					buf.Write(text)
// 					pendingSelector = false
// 				} else if seenColon {
// 					// This is likely a property value
// 					buf.Write(text)
// 					buf.WriteByte(' ')
// 				} else {
// 					// This is a top-level selector
// 					if startOfLine {
// 						buf.WriteString(currentIndent)
// 					}
// 					buf.Write(text)
// 					buf.WriteByte(' ')
// 					pendingSelector = true
// 				}
// 				startOfLine = false
// 			} else {
// 				// Minification mode
// 				buf.Write(text)
// 			}

// 		case css.DelimToken:
// 			// Handle delimiters, including the dot for class selectors
// 			if string(text) == "." {
// 				if params.isBeautifier && startOfLine && inRuleset {
// 					buf.WriteString(currentIndent)
// 					startOfLine = false
// 				}
// 				buf.WriteByte('.')
// 				pendingSelector = true
// 			} else {
// 				buf.Write(text)
// 			}

// 		case css.ColonToken:
// 			seenColon = true
// 			if params.isBeautifier {
// 				buf.WriteString(": ")
// 				startOfLine = false
// 			} else {
// 				buf.WriteByte(':')
// 			}

// 		case css.SemicolonToken:
// 			if params.isBeautifier {
// 				buf.WriteByte(';')
// 				buf.WriteByte('\n')
// 				seenColon = false
// 				startOfLine = true
// 				pendingSelector = false
// 			} else {
// 				// For minification, add semicolon
// 				buf.WriteByte(';')
// 				seenColon = false
// 			}

// 		case css.LeftBraceToken:
// 			inRuleset = true
// 			pendingSelector = false

// 			if params.isBeautifier {
// 				buf.WriteByte('{')
// 				buf.WriteByte('\n')
// 				indentLevel++
// 				startOfLine = true
// 			} else {
// 				buf.WriteByte('{')
// 			}

// 		case css.RightBraceToken:
// 			if params.isBeautifier {
// 				if indentLevel > 0 {
// 					indentLevel--
// 				}
// 				currentIndent = strings.Repeat(nestedIndentUnit, indentLevel)
// 				buf.WriteString(currentIndent)
// 				buf.WriteByte('}')
// 				buf.WriteByte('\n')
// 				startOfLine = true
// 			} else {
// 				// For minification: remove the last semicolon if present
// 				// Get current buffer content
// 				bufContent := buf.Bytes()

// 				// Check if the last character is a semicolon
// 				if len(bufContent) > 0 && bufContent[len(bufContent)-1] == ';' {
// 					// Remove only the semicolon
// 					buf.Truncate(buf.Len() - 1)
// 				}

// 				buf.WriteByte('}')
// 			}

// 			if indentLevel == 0 {
// 				inRuleset = false
// 			}

// 		case css.NumberToken, css.PercentageToken, css.DimensionToken:
// 			// Handle numbers and dimensions (like px, em, etc.)
// 			buf.Write(text)
// 			if params.isBeautifier {
// 				startOfLine = false
// 			}

// 		case css.CommaToken:
// 			if params.isBeautifier {
// 				buf.WriteByte(',')
// 				buf.WriteByte(' ')
// 				startOfLine = false
// 			} else {
// 				buf.WriteByte(',')
// 			}

// 		case css.WhitespaceToken:
// 			if params.isBeautifier {
// 				// Only add a space if not at start of line and not after certain tokens
// 				if !startOfLine && !pendingSelector {
// 					buf.WriteByte(' ')
// 				}
// 			}
// 			// In minification mode, skip all whitespace

// 		case css.HashToken:
// 			// Handle ID selectors
// 			if params.isBeautifier {
// 				if startOfLine {
// 					buf.WriteString(currentIndent)
// 				}
// 				buf.Write(text)
// 				buf.WriteByte(' ')
// 				startOfLine = false
// 				pendingSelector = true
// 			} else {
// 				buf.Write(text)
// 			}

// 		case css.AtKeywordToken:
// 			// Handle at-rules like @media
// 			if params.isBeautifier {
// 				buf.WriteString(currentIndent)
// 				buf.Write(text)
// 				buf.WriteByte(' ')
// 				startOfLine = false
// 			} else {
// 				buf.Write(text)
// 			}

// 		case css.CommentToken:
// 			if params.isBeautifier {
// 				buf.WriteString(currentIndent)
// 				buf.Write(text)
// 				buf.WriteByte('\n')
// 				startOfLine = true
// 			}
// 			// Skip comments in minification mode

// 		default:
// 			// Handle other tokens
// 			buf.Write(text)
// 			if params.isBeautifier {
// 				startOfLine = false
// 			}
// 		}
// 	}

// 	return buf.String(), nil
// }

func BeautifyCSS(params *ParamsUpdateCSS) (string, error) {
	return normalizeCSS(
		&paramsNormalizeCSS{
			ParamsUpdateCSS: *params,
			isBeautifier:    true,
		},
	)
}

func MinifyCSS(params *ParamsUpdateCSS) (string, error) {
	return normalizeCSS(
		&paramsNormalizeCSS{
			ParamsUpdateCSS: *params,
		},
	)
}
