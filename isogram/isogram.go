package isogram

import (
	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
	"strings"
)

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	regex := pcre.MustCompile(`(\w).*\1`, 0)

	return !regex.MatcherString(word, 0).Matches()
}