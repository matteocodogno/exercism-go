package acronym

import (
	"regexp"
	"strings"
)

func Abbreviate(s string) string {
	re := regexp.MustCompile(`[-_]+`)
	s = re.ReplaceAllString(s, " ")
	words := strings.Fields(s)
	acronym := ""
	for _, word := range words {
		acronym = acronym + string(word[0])
	}

	return strings.ToUpper(acronym)
}
