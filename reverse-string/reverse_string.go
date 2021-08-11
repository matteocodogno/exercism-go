package reverse

import "unicode/utf8"

func Reverse(word string) string {
	reverse := ""
	for len(word) > 0 {
		r, size := utf8.DecodeLastRuneInString(word)
		reverse += string(r)

		word = word[:len(word)-size]
	}

	return reverse
}