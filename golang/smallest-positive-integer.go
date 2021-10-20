package main

// This is a demo task.
//
//Write a function:
//
//func Solution(A []int) int
//
//that, given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.
//
//For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.
//
//Given A = [1, 2, 3], the function should return 4.
//
//Given A = [−1, −3], the function should return 1.
//
//Write an efficient algorithm for the following assumptions:
//
//N is an integer within the range [1..100,000];
//each element of array A is an integer within the range [−1,000,000..1,000,000].

func smallestPositiveInteger(nums []int) int {
	keys := make(map[int]int, len(nums))

	isMissing := false
	missing := 1

	for _, num := range nums {
		if num > 0 {
			keys[num] = num
		}
	}

	for !isMissing {
		if _, ok := keys[missing]; !ok {
			isMissing = true
			break
		}

		missing++
	}

	return missing
}
