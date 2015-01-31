--[[
The first two consecutive numbers to have two distinct prime factors are:

14 = 2 × 7
15 = 3 × 5

The first three consecutive numbers to have three distinct prime factors are:

644 = 2² × 7 × 23
645 = 3 × 5 × 43
646 = 2 × 17 × 19.

Find the first four consecutive integers to have four distinct prime factors. What is the first of these numbers?
]]

require "utils/primes"

primeSet, primeArray = primes.upto(1000)

function factors (n)
	local factors = {}
	local r = n

	for _, p in ipairs(primeArray) do
		if r == 1 or p == n then 
			break
		end

		if r % p == 0 then
			table.insert(factors, p)

			while r % p == 0 do
				r = r / p
			end
		end
	end

	return factors
end

consecutive = {}
i = 210 -- start at 2 * 3 * 5 * 7
while #consecutive < 4 do
	if #factors(i) == 4 then
		table.insert(consecutive, i)
	else
		consecutive = {}
	end

	i = i + 1
end

for _, int in ipairs(consecutive) do print(int) end

