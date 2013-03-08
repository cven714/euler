package utils

func PrimesUpTo(n uint64) []uint64 {
	ints := make([]bool, n)

	primes := make([]uint64, 0, n/4)
	primes = append(primes, 2)

	var i, j uint64

	for i = 3; i < n; i += 2 {
		if !ints[i] {
			primes = append(primes, i)
			for j = i; j < n; j += i {
				ints[j] = true
			}
		}
	}

	return primes
}

func MapPrimesUpTo(n uint64) map[uint64]struct{} {
	ints := make([]bool, n)
	prime := struct{}{}

	primes := make(map[uint64]struct{}, n/4)
	primes[2] = prime

	var i, j uint64

	for i = 3; i < n; i += 2 {
		if !ints[i] {
			primes[i] = prime
			for j = i; j < n; j += i {
				ints[j] = true
			}
		}
	}

	return primes
}