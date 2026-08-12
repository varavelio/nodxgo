package nodx

import "strings"

// escapeChars are the characters that must be escaped to prevent XSS.
const escapeChars = "&<>\"'"

// EscapeHTML escapes the input string to prevent XSS attacks.
//
// If the input contains no characters that need escaping, it is returned
// unchanged without any allocation.
func EscapeHTML(input string) string {
	// Fast path: nothing to escape, return the input as-is.
	if !strings.ContainsAny(input, escapeChars) {
		return input
	}

	var builder strings.Builder
	builder.Grow(len(input))
	start := 0

	for i := 0; i < len(input); i++ {
		var replacement string
		switch input[i] {
		case '&':
			replacement = "&amp;"
		case '<':
			replacement = "&lt;"
		case '>':
			replacement = "&gt;"
		case '"':
			replacement = "&quot;"
		case '\'':
			replacement = "&#39;"
		default:
			continue
		}
		builder.WriteString(input[start:i])
		builder.WriteString(replacement)
		start = i + 1
	}

	builder.WriteString(input[start:])
	return builder.String()
}
