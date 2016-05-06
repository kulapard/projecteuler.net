# -*- coding: utf-8 -*-
"""
Special Pythagorean triplet
https://projecteuler.net/problem=9
"""
import math


def abc_gen(limit):
    for a in xrange(limit):
        for b in xrange(a + 1, limit):
            c = math.sqrt(a ** 2 + b ** 2)
            if c == int(c) and a < b < c:
                yield a, b, int(c)


for a, b, c in abc_gen(1000):
    if a + b + c == 1000:
        print a * b * c
