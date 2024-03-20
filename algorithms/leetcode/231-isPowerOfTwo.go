// Given an integer n, return true if it is a power of two. Otherwise, return false.

// An integer n is a power of two, if there exists an integer x such that n == 2x.

// Example 1:

// Input: n = 1
// Output: true
// Explanation: 2^0 = 1
// Example 2:

// Input: n = 16
// Output: true
// Explanation: 2^4 = 16
// Example 3:

// Input: n = 3
// Output: false

package main

func isPowerOfTwo(n int) bool {
	// To be a power of two, a number must be greater than 0
	// and have only one bit set to 1 in its binary representation.
	// For example, 2 (binary: 10), 4 (binary: 100), 8 (binary: 1000), etc.

	// Check if n is greater than 0 and if the bitwise AND operation
	// between n and (n - 1) results in 0.
	// (n - 1) unsets the rightmost 1 bit and sets all the bits to the right of it.
	// So, if n is a power of two, n & (n - 1) will result in 0.
	return n > 0 && n&(n-1) == 0
}
