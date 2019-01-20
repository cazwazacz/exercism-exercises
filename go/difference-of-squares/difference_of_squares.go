package diffsquares

import "math"

// SquareOfSum returns the square of the sum of the first n natural numbers
func SquareOfSum(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		sum += i
	}

	return squareAsInt(sum)
}

// SumOfSquares returns the sum of the squares of the first n natural numbers
func SumOfSquares(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		sum += squareAsInt(i)
	}

	return sum
}

// Difference returns the difference between square of the sum and the sum of the squares of the first n natural numbers
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

const square float64 = 2

func squareAsInt(number int) int {
	squared := math.Pow(float64(number), square)

	return int(squared)
}
