# -*- coding: utf-8 -*-
import math


def is_prime(n):
    for i in xrange(2, int(math.sqrt(n)) + 1):
        if n % i == 0:
            return False

    return True
