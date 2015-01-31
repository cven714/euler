primes = {}

function primes.upto (n)
	if n < 2 then return end

	local primeSet = { ["2"] = true }
	local primeArray = { 2 }
	local composites = {}

	for i = 3, n + 1, 2 do
		if composites[i] == nil then
			primeSet[i] = true
			table.insert(primeArray, i)

			for j = i, n + 1, i do
				composites[j] = true
			end
		end
	end

	return primeSet, primeArray
end