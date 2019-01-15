package raindrops

import "strconv"

// Convert takes an integer and returns a string depending on which factors the numbers have
// if the number has 3 as a factor, output "Pling"
// if the number has 5 as a factor, output "Plang"
// if the number has 7 as a factor, output "Plong"
// if the number does not have any of the numbers as a factor, it gets passed through as a string
func Convert(number int) string {
	resultString := ""

	if number%3 == 0 {
		resultString += "Pling"
	}

	if number%5 == 0 {
		resultString += "Plang"
	}

	if number%7 == 0 {
		resultString += "Plong"
	}

	if resultString == "" {
		return strconv.Itoa(number)
	}

	return resultString
}
