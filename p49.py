# The arithmetic sequence, 1487, 4817, 8147, in which each of the terms increases by 3330, is unusual in two ways: (i) each of the three terms are prime, and, (ii) each of the 4-digit numbers are permutations of one another.
#
# There are no arithmetic sequences made up of three 1-, 2-, or 3-digit primes, exhibiting this property, but there is one other 4-digit increasing sequence.
#
# What 12-digit number do you form by concatenating the three terms in this sequence?

from imp import load_source
primes = load_source('primes', 'utils/primes.py')

def permutations(x, y, z):
  return sorted(str(x)) == sorted(str(y)) == sorted(str(z))

prime_set = { p for p in primes.upto(10000) if p > 1000 }

for prime in prime_set:
  x, y = prime + 3330, prime + 6660

  if x in prime_set and y in prime_set and permutations(prime, x, y):
    print prime, x, y
