package main

// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

// Input: nums = [2,2,1] Output: 1
// Input: nums = [4,1,2,1,2] Output: 4
// Input: nums = [1] Output: 1

// Constraints:
// => 1 <= nums.length <= 3 * 10^4
// => -3 * 104 <= nums[i] <= 3 * 104
// Each element in the array appears twice except for one element which appears only once.

func singleNumber(nums []int) int {
	values := make(map[int]int)
	for _, num := range nums {
		values[num] += 1
	}

	for i, v := range values {
		if v == 1 {
			return i
		}
	}

	return 1

	//var single int
	//for _, n := range nums {
	//	single ^= n
	//}
	//
	//return single
}
