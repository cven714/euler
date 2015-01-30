--[[
The series, 1^1 + 2^2 + 3^3 + ... + 10^10 = 10405071317.

Find the last ten digits of the series, 1^1 + 2^2 + 3^3 + ... + 1000^1000.
]]

sum = 0
for i = 1, 1000 do
	sum = (sum + i^i) % 10^10 -- Would work if I had an arbitrary precision number lib available.
end

print(string.format("%.0f", sum))