# The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

# Find the sum of all the primes below two million.

import imp
primes = imp.load_source('primes', 'utils/primes.py')

print sum(primes.upto(2000000))