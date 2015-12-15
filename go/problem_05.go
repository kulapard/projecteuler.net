/*
Smallest multiple
https://projecteuler.net/problem=5

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package main

import (
	"fmt"
)

func isDivisibleByAllInRange(num, min, max int) bool {
	for i := min; i <= max; i++ {
		if num%i != 0 {
			return false
		}
	}
	return true
}

func main() {
	min := 1
	max := 20
	fmt.Printf("min = %d, max = %d\n", min, max)

	num := max
	for !isDivisibleByAllInRange(num, min, max) {
		num++
	}
	fmt.Printf("Bingo %d\n", num)
}
