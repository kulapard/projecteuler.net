# -*- coding: utf-8 -*-
"""
Gozinta Chains
https://projecteuler.net/problem=548
"""


def divisors(a, b):
    for i in xrange(a + 1, b):
        if i % a == 0 and b % i == 0:
            yield i


def gozinta_chains(a, b):
    if a == b or b == 0:
        return

    yield [a, b]

    for d in divisors(a, b):
        for sub_chain in gozinta_chains(d, b):
            yield [a] + sub_chain


def g(n):
    count = 0
    for _ in gozinta_chains(1, n):
        count += 1
    return count

# result = 0
# for n in xrange(10 ** 16):
#     if g(n) == n:
#         print n
#         result += n
#
#     if n % 1000 == 0:
#         print n
#
# print result

for n in xrange(100):
    print n, g(n)
