// Package diffsquares for finding the difference between square of sum and sum of square
package algorithms

// SquareOfSum function calculates the square of sum
func SquareOfSum(n int) int {
	if n <= 0 {
        return 0
    }
    res := int((n * (n + 1)) / 2)
    return res * res
}

// SumOfSquares function calculates the sum of squares
func SumOfSquares(n int) int {
	if n <= 0 {
        return 0
    }
    sum := int(((n *(n + 1) * (2*n + 1)) / 6))
    return sum
}

// Difference consolidates the result
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
