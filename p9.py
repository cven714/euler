# A Pythagorean triplet is a set of three natural numbers, a < b < c, 
# for which a^2 + b^2 = c^2

# For example, 32 + 42 = 9 + 16 = 25 = 52.

# There exists exactly one Pythagorean triplet for which a + b + c = 1000.
# Find the product abc.

for m in xrange(30, 1, -1):
	for n in xrange(m - 1, 0, -1):
		a, b, c = (m * m) - (n * n), 2 * m * n, (m * m) + (n * n)
		if a + b + c == 1000:
			print a * b * c