package main

// Given an array nums of n integers where nums[i] is in the range [1, n],
// return an array of all the integers in the range [1, n] that do not appear in nums.

// Input: nums = [4,3,2,7,8,2,3,1]  Output: [5,6]
// Input: nums = [1,1]  Output: [2]

// Constraints:
// * n == nums.length
// * 1 <= n <= 10^5
// * 1 <= nums[i] <= n

func findDisappearedNumbers(nums []int) []int {
	var values []int
	indexes := make([]int, len(nums)+1)

	for _, num := range nums {
		indexes[num]++
	}

	for i := 1; i < len(indexes); i++ {
		if indexes[i] == 0 {
			values = append(values, i)
		}
	}

	return values
}
