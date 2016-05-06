# -*- coding: utf-8 -*-
"""
Summation of primes
https://projecteuler.net/problem=10
"""
from utils import is_prime


def prime_gen():
    x = 2
    while True:
        if is_prime(x):
            yield x
        x += 1


prime_sum = 0
next_prime = 0
limit = 2 * 10 ** 6
prime = prime_gen()
while next_prime < limit:
    prime_sum += next_prime
    next_prime = next(prime)

print prime_sum
