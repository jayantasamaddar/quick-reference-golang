// Given an integer array nums, return the greatest common divisor of the smallest number and largest number in nums.

// The greatest common divisor of two numbers is the largest positive integer that evenly divides both numbers.

// Example 1:

// Input: nums = [2,5,6,9,10]
// Output: 2
// Explanation:
// The smallest number in nums is 2.
// The largest number in nums is 10.
// The greatest common divisor of 2 and 10 is 2.
// Example 2:

// Input: nums = [7,5,6,8,3]
// Output: 1
// Explanation:
// The smallest number in nums is 3.
// The largest number in nums is 8.
// The greatest common divisor of 3 and 8 is 1.
// Example 3:

// Input: nums = [3,3]
// Output: 3
// Explanation:
// The smallest number in nums is 3.
// The largest number in nums is 3.
// The greatest common divisor of 3 and 3 is 3.

package main

// Solution:
//
// We need the min and the max of the list.
// From the larger of the two numbers we should loop backwards towards 2 (i > 1) and on finding the first match return the number
// The match is max % i == 0 && min % i == 0. i.e. The number should evenly divide both the max and the min
// The loop terminates before i == 1, and since 1 divides both the min and max, this is our default return value
func minMax(nums []int) (int, int) {
	var min, max = nums[0], nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	return min, max
}

func findGCD(nums []int) int {
	min, max := minMax(nums)

	for i := max; i > 1; i-- {
		if max%i == 0 && min%i == 0 {
			return i
		}
	}
	return 1
}
