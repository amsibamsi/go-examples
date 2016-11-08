package revstr

import "bytes"

func ReverseString(s string) string {
	chars := []rune(s)
	var buf bytes.Buffer
	for i := len(chars) - 1; i >= 0; i-- {
		buf.WriteRune(chars[i])
	}
	return buf.String()
}
