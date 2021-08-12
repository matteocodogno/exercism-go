package luhn

import (
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func reverse(word string) string {
	reverse := ""
	for len(word) > 0 {
		r, size := utf8.DecodeLastRuneInString(word)
		reverse += string(r)

		word = word[:len(word)-size]
	}

	return reverse
}

func Valid(s string) bool {
	checksum := 0
	s = strings.ReplaceAll(s, " ", "")
	regex := regexp.MustCompile(`^\d{2,}$`)

	if !regex.MatchString(s) {
		return false
	}

	s = reverse(s)
	for i := range s {
		number, _ := strconv.Atoi(string(s[i]))
		if i % 2 == 1 {
			doubled := number * 2
			if doubled > 9 {
				doubled -= 9
			}
			checksum += doubled
		} else {
			checksum += number
		}
	}

	return checksum % 10 == 0
}