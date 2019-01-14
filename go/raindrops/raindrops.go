package raindrops

import (
	"strconv"
	"strings"
)

type factor struct {
	number int
	string string
}

var factors = []factor{
	factor{3, "Pling"},
	factor{5, "Plang"},
	factor{7, "Plong"},
}

// Convert takes an integer and returns a string depending on which factors the numbers have
// if the number has 3 as a factor, output "Pling"
// if the number has 5 as a factor, output "Plang"
// if the number has 7 as a factor, output "Plong"
// if the number does not have any of the numbers as a factor, it gets passed through as a string
func Convert(number int) string {
	var stringBuilder strings.Builder

	for _, factor := range factors {
		stringBuilder.WriteString(checkFactor(number, factor.number, factor.string))
	}

	if stringBuilder.String() == "" {
		return strconv.Itoa(number)
	}

	return stringBuilder.String()
}

func checkFactor(number int, factor int, returnString string) string {
	if number%factor == 0 {
		return returnString
	}

	return ""
}
