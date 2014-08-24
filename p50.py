# The prime 41, can be written as the sum of six consecutive primes:

# 41 = 2 + 3 + 5 + 7 + 11 + 13
# This is the longest sum of consecutive primes that adds to a prime below one-hundred.

# The longest sum of consecutive primes below one-thousand that adds to a prime, contains 21 terms, and is equal to 953.

# Which prime, below one-million, can be written as the sum of the most consecutive primes?

import imp
primes = imp.load_source('primes', 'utils/primes.py')

mil = 1000000
prime_list = list(primes.upto(mil))
prime_set = set(prime_list)

consecutive = 22
prime = 0
n = len(prime_list)

for i in range(n):
  c = consecutive
  s = sum(prime_list[i:i+c])
  while s < mil and i + c < n:
    if s in prime_set:
      consecutive, prime = c, s

    s += prime_list[i + c]
    c += 1

print consecutive, prime
