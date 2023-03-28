package env

import (
	"bytes"
)

func EscapeSpecialChars(value string) string {
	var buf bytes.Buffer
	for _, r := range value {
		switch r {
		case '\n':
			buf.WriteString("\\n")
		case '\r':
			buf.WriteString("\\r")
		case '\t':
			buf.WriteString("\\t")
		default:
			buf.WriteRune(r)
		}
	}
	return buf.String()
}
