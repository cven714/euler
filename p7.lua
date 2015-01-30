--[[
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, 
we can see that the 6th prime is 13.

What is the 10 001st prime number?
]]

primes = { 2, 3, 5, 7, 11, 13 }
i = 15

while #primes < 10001 do
	i = i + 2
	for _, p in ipairs(primes) do
		if i % p == 0 then break end

		if p > i / 4 then 
			table.insert(primes, i)
			break 
		end
	end
end

print(primes[#primes])
