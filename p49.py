import math
import functools

def arr_max(arr):
    return functools.reduce(lambda acc, x: x if x > acc else acc, arr, arr[0])

def is_permutation(a, b):
    result = True
    sa, sb = [x for x in str(a)], [x for x in str(b)]
    sa.sort()
    sb.sort()

    for i in range(len(sa)):
        if sa[i] != sb[i]:
            result = False
            break

    return result

def is_prime(n, primes = None):
    if primes == None:
        primes = range(1, math.sqrt(n) + 1)

    result = True
    for i in primes:
        if n % i == 0:
            result = False
            break
    return result

def eratosthene(n):
    if n < 2:
        return []

    primes = [2]
    for i in range(3, n + 1, 2):
        if (is_prime(i, primes)):
            primes.append(i)

    return primes

primes = [x for x in eratosthene(9999) if x >= 1000]

permutations = []

for index, value in enumerate(primes):
    permutation = []
    for i in range(index + 1, len(primes)):
        if is_permutation(value, primes[i]):
            permutation.append(primes[i])

    if len(permutation) >= 2:
        permutation.append(value)
        permutations.append(permutation)

for v in permutations:
    result = []
    for i in v:
        answer = 0
        for j in v:
            if j - i == 3330:
                answer = i
                break
        if answer != 0:
            result.append(answer)

    if len(result) >= 2:
        result.append(arr_max(result) + 3330)
        print(result)