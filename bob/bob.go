package bob

import (
	"regexp"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	answer := "Whatever."

	remark = strings.TrimRight(remark, " ")
	isYelling, _ := regexp.MatchString(`^[^a-z]*[A-Z][^a-z]*[^\?]$`, remark)
	isYellingAQuestion, _ := regexp.MatchString(`^[A-Z0-9',\s!]*\?$`, remark)
	isAQuestion, _ := regexp.MatchString(`^[A-Za-z0-9,:\(\)\s!\.]*\?$`, remark)
	isSilence, _ := regexp.MatchString(`^[\s]*$`, remark)

	switch {
	case isYelling:
		answer = "Whoa, chill out!"
	case isAQuestion:
		answer = "Sure."
	case isYellingAQuestion:
		answer = "Calm down, I know what I'm doing!"
	case isSilence:
		answer = "Fine. Be that way!"
	}

	return answer
}
