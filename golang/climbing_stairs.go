package main

// You are climbing a staircase. It takes n steps to reach the top.

// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

// Input: n = 2 Output: 2
// Explanation: There are two ways to climb to the top. 1 step + 1 step, 2 steps

// Input: n = 3
// Output: 3
// Explanation: There are three ways to climb to the top.
// 1 step + 1 step + 1 step, 1 step + 2 steps, 2 steps + 1 step

var cache = map[int]int{}

func climbStairs(n int) int {
	if n < 3 {
		return n
	}

	if cache[n] > 0 {
		return cache[n]
	}

	cache[n] = climbStairs(n-1) + climbStairs(n-2)
	return cache[n]
}
