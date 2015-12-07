/*
Largest prime factor
https://projecteuler.net/problem=3

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

package main

import "fmt"

func main() {
	num := 600851475143

	for {
		divider := 2
		for num%divider != 0 {
			divider++
		}

		if divider == num || divider == 1 {
			fmt.Printf("Bingo! %d is a prime factor!\n", num)
			break
		}

		fmt.Printf("%d = %d * %d\n", num, num/divider, divider)
		num = num / divider
	}

}
