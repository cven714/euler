# A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 x 99.

# Find the largest palindrome made from the product of two 3-digit numbers.

def palindrome(s):
	for i in range(0, len(s)/2):
		if s[i] != s[-(i+1)]: return False

	return True

maxPalindrome = 0

for i in range(999, 99, -1):
	for j in range(i, 99, -1):
		p = i * j
		if palindrome(str(p)) and p > maxPalindrome:
			maxPalindrome = p

print maxPalindrome