import time, itertools

def generate():
	yield 2
	yield 3
	yield 5
	yield 7
	yield 11
	yield 13

	primes = [2, 3, 5, 7, 11, 13]

	for i in itertools.count(17, 2):
		if all(i % p for p in primes if p < i / 4):
			yield i
			primes.append(i)

def upto(n):
	if n < 2:
		return
	else:
		yield 2
	
	max = n + 1
	prime = [i % 2 != 0 for i in xrange(max)]
	prime[1], prime[2] = False, True

	for i in xrange(3, max, 2):
		if prime[i]:
			yield i
			for m in xrange(i + i, max, i):
				prime[m] = False

def nth(n):
	return list(itertools.islice(generate(), n - 1, n))[0]