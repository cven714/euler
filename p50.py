# The prime 41, can be written as the sum of six consecutive primes:

# 41 = 2 + 3 + 5 + 7 + 11 + 13
# This is the longest sum of consecutive primes that adds to a prime below one-hundred.

# The longest sum of consecutive primes below one-thousand that adds to a prime, contains 21 terms, and is equal to 953.

# Which prime, below one-million, can be written as the sum of the most consecutive primes?

import imp
primes = imp.load_source('primes', 'utils/primes.py')

mil = 1000000
primeList = list(primes.upto(mil))
primeSet = set(primeList)

consecutive = 22
prime = 0

for i in range(len(primeList)):
  print i
  c = consecutive
  s = sum(primeList[i:c])
  while s < mil and i + c < len(primeList):
    if s in primeSet:
      #print c, s
      consecutive, prime = c, s

    s += primeList[i + c]
    c += 1

print consecutive, prime


