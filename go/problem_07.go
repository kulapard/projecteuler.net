/*
10001st prime
https://projecteuler.net/problem=7

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
*/

package main

import (
	"fmt"
	"math"
	"time"
)

func isPrime(n int) bool {
	max := int(math.Sqrt(float64(n)))
	for i := 2; i <= max; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Sieve of Eratosthenes
// https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
func buildSieveOfEratosthenes(a []bool, n int) []bool {
	//	println("BuildSieveOfEratosthenes")
	newA := make([]bool, n-len(a))
	for i := range newA {
		newA[i] = true
	}

	a = append(a, newA...)

	for i := 2; int(math.Pow(float64(i), 2)) <= n; i++ {
		if a[i] {
			for j := int(math.Pow(float64(i), 2)); j < n; j = j + i {
				a[j] = false
			}
		}
	}
	return a
}

func bruteForce(indexToFind int) int {
	prime := 0
	primeCounter := 0
	for i := 2; primeCounter < indexToFind; i++ {
		if isPrime(i) {
			primeCounter++
			prime = i
		}
	}
	return prime
}

func sieveOfEratosthenes(indexToFind int) int {
	var prime int

	primeCounter := 0
	sieveSize := indexToFind * 20
	sieve := []bool{}
	for primeCounter < indexToFind {
		primeCounter = 0
		sieve := buildSieveOfEratosthenes(sieve, sieveSize)

		prime = 0
		//	prime_counter = 0
		for i := 2; i < len(sieve); i++ {
			if sieve[i] {
				primeCounter++
				prime = i
				if primeCounter == indexToFind {
					break
				}
			}
		}
		sieveSize = sieveSize * 2
	}
	return prime
}

func main() {
	indexToFind := 10001

	var prime int
	var t time.Time

	// 1) Brute force (less memory)
	t = time.Now()
	prime = bruteForce(indexToFind)
	fmt.Printf("[%s] Brute force: The %dst prime number is %d\n", time.Since(t), indexToFind, prime)

	// 2) Sieve of Eratosthenes (faster, sieve size)
	t = time.Now()
	prime = sieveOfEratosthenes(indexToFind)
	fmt.Printf("[%s] Sieve of Eratosthenes: The %dst prime number is %d\n", time.Since(t), indexToFind, prime)
}
