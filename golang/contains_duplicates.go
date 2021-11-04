package main

import (
	"sort"
)

// Given an integer array nums, return true if any value appears at least twice in the array,
// and return false if every element is distinct.

// Example 1:
// Input: nums = [1,2,3,1]
// Output: true

// Example 2:
// Input: nums = [1,2,3,4]
// Output: false

// Example 3:
// Input: nums = [1,1,1,3,3,4,3,2,4,2]
// Output: true

func containsDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	sort.Ints(nums)

	pairs := make(map[int]int)

	for _, num := range nums {
		if _, ok := pairs[num]; ok {
			return true
		}

		pairs[num] = pairs[num]
	}

	return false
}
