/*
Largest palindrome product
https://projecteuler.net/problem=4

A palindromic number reads the same both ways. The largest palindrome made from the product
of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isPalindromic(num int) bool {
	numStr := strconv.Itoa(num)
	return numStr == reverse(numStr)
}

func hasDividerInRange(num, min, max int) bool {
	for i := min; i <= max; i++ {
		k := num / i
		if num%i == 0 && k >= min && k <= max {
			fmt.Printf("%d = %d * %d\n", num, i, k)
			return true
		}
	}
	return false
}

func main() {
	digits := 3
	min := int(math.Pow(10, float64(digits-1)))
	max := int(math.Pow(10, float64(digits)) - 1)
	fmt.Printf("min = %d, max = %d\n", min, max)

	for num := max * max; num >= min; num-- {
		if isPalindromic(num) && hasDividerInRange(num, min, max) {
			fmt.Printf("Bingo! %d\n", num)
			break
		}
	}

}
