package scrabble

import "strings"

func Score(word string) int {
	score := 0

	for i := range word {
		switch strings.ToUpper(string(word[i])) {
		case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T":
			score += 1
		case "D", "G":
			score += 2
		case "B", "C", "M", "P":
			score += 3
		case "F", "H", "V", "W", "Y":
			score += 4
		case "K":
			score += 5
		case "J", "X":
			score += 8
		case "Q", "Z":
			score += 10
		}
	}

	return score
}