package scrabble

import "strings"

// Score takes a string and returns its scrabble score
func Score(word string) int {
	score := 0

	for _, character := range word {
		for points, letters := range lettersToPoints {
			if letterInSlice(character, letters) {
				score += points
			}
		}
	}

	return score
}

func letterInSlice(character rune, slice []string) bool {
	for i := range slice {
		if strings.ToUpper(string(character)) == slice[i] {
			return true
		}
	}

	return false
}

var lettersToPoints = map[int][]string{
	1:  {"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
	2:  {"D", "G"},
	3:  {"B", "C", "M", "P"},
	4:  {"F", "H", "V", "W", "Y"},
	5:  {"K"},
	8:  {"J", "X"},
	10: {"Q", "Z"},
}
