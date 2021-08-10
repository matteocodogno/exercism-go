package raindrops

import "strconv"

func Convert(raindrops int) string {
	result := ""

	if raindrops % 3 == 0 {
		result = "Pling"
	}

	if raindrops % 5 == 0 {
		result = result + "Plang"
	}

	if raindrops % 7 == 0 {
		result = result + "Plong"
	}

	if result == "" {
		result = strconv.Itoa(raindrops)
	}

	return result
}