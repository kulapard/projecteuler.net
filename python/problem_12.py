# -*- coding: utf-8 -*-
"""
Highly divisible triangular number
https://projecteuler.net/problem=12
"""
import math


def triangle_numbers():
    x = t = 1
    while True:
        yield t
        x += 1
        t += x


def divisors(n):
    for i in xrange(1, int(math.sqrt(n)) + 1):
        if n % i == 0:
            yield i


def divisors_count(n):
    count = 0
    for _ in divisors(n):
        count += 1
    return count * 2 if n > 1 else 1


for triangle_number in triangle_numbers():
    if divisors_count(triangle_number) > 500:
        print triangle_number
        break
