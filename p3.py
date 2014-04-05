# The prime factors of 13195 are 5, 7, 13 and 29.

# What is the largest prime factor of the number 600851475143 ?

def primeFactors(n):
    factors = []

    while n > 1 and n % 2 == 0:
        factors.append(2)
        n /= 2

    d = 3
    while n > 1:
        while n % d == 0:
            factors.append(d)
            n /= d
        d += 2
        if d*d > n:
            if n > 1: factors.append(n)
            break
    return factors

print max(primeFactors(600851475143))