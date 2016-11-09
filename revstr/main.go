package revstr

import "bytes"

func ReverseStringBuffer(s string) string {
	r := []rune(s)
	var b bytes.Buffer
	for i := len(r) - 1; i >= 0; i-- {
		b.WriteRune(r[i])
	}
	return b.String()
}

func ReverseStringAppend(s string) string {
	r := []rune(s)
	t := []rune{}
	for i := len(r) - 1; i >= 0; i-- {
		t = append(t, r[i])
	}
	return string(t)
}

func ReverseStringIndex(s string) string {
	r := []rune(s)
	var t []rune
	t = make([]rune, len(r))
	for i := 0; i < len(r); i++ {
		t[i] = r[len(r)-1-i]
	}
	return string(t)
}

func ReverseStringInplace(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
