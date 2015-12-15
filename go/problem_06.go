/*
Sum square difference
https://projecteuler.net/problem=6

The sum of the squares of the first ten natural numbers is,
1^2 + 2^2 + ... + 10^2 = 385

The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)^2 = 55^2 = 3025

Hence the difference between the sum of the squares of the first ten natural numbers
and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

package main

import (
	"fmt"
	"math"
)

func squareOfSum(n int) int {
	// Sum of arithmetic progression
	sum := n * (1 + n) / 2
	return int(math.Pow(float64(sum), 2))
}

func sumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

func main() {
	num := 100
	squareOfSum := squareOfSum(num)
	sumOfSquares := sumOfSquares(num)
	fmt.Printf("squareOfSum: %d\n", squareOfSum)
	fmt.Printf("sumOfSquares: %d\n", sumOfSquares)
	fmt.Printf("Bingo %d\n", squareOfSum-sumOfSquares)
}
