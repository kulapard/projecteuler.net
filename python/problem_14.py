# -*- coding: utf-8 -*-
"""
Longest Collatz sequence
https://projecteuler.net/problem=14
"""


def collatz_sequence(n):
    yield n

    while n > 1:
        if n % 2 == 0:
            n /= 2
        else:
            n = 3 * n + 1

        yield n


def collatz_sequence_size(n):
    count = 0
    for _ in collatz_sequence(n):
        count += 1
    return count


max_start_number = 0
max_chain_size = 0
for start_number in xrange(10 ** 6):
    curr_chain_size = collatz_sequence_size(start_number)
    if curr_chain_size > max_chain_size:
        max_start_number = start_number
        max_chain_size = curr_chain_size

print max_start_number, max_chain_size
