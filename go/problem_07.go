/*
10001st prime
https://projecteuler.net/problem=7

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
*/

package main

import (
	"math"
	"fmt"
	"time"
)

func isPrime(n int) bool {
	max := int(math.Sqrt(float64(n)))
	for i := 2; i <= max; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

// Sieve of Eratosthenes
// https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
func BuildSieveOfEratosthenes(a []bool, n int, ) []bool {
//	println("BuildSieveOfEratosthenes")
	new_a := make([]bool, n - len(a))
	for i := range new_a {
		new_a[i] = true
	}

	a = append(a, new_a...)


	for i := 2; int(math.Pow(float64(i), 2)) <= n; i++ {
		if a[i] {
			for j := int(math.Pow(float64(i), 2)); j < n; j = j + i {
				a[j] = false
			}
		}
	}
	return a
}

func BruteForce(index_to_find int) int {
	prime := 0
	prime_counter := 0
	for i := 2; prime_counter < index_to_find; i++ {
		if isPrime(i) {
			prime_counter++
			prime = i
		}
	}
	return prime
}

func SieveOfEratosthenes(index_to_find int) int {
	var prime int

	prime_counter := 0
	sieve_size := index_to_find * 20
	sieve := []bool{}
	for prime_counter < index_to_find {
		prime_counter = 0
		sieve := BuildSieveOfEratosthenes(sieve, sieve_size)

		prime = 0
		//	prime_counter = 0
		for i := 2; i < len(sieve); i++ {
			if sieve[i] {
				prime_counter++
				prime = i
				if prime_counter == index_to_find {
					break
				}
			}
		}
		sieve_size = sieve_size * 2
	}
	return prime
}

func main() {
	index_to_find := 10001

	var prime int
	var t time.Time

	// 1) Brute force (less memory)
	t = time.Now()
	prime = BruteForce(index_to_find)
	fmt.Printf("[%s] Brute force: The %dst prime number is %d\n", time.Since(t), index_to_find, prime)

	// 2) Sieve of Eratosthenes (faster, sieve size)
	t = time.Now()
	prime = SieveOfEratosthenes(index_to_find)
	fmt.Printf("[%s] Sieve of Eratosthenes: The %dst prime number is %d\n", time.Since(t), index_to_find, prime)
}
