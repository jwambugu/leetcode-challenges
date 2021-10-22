package main

// Given a string s, find the length of the longest substring without repeating characters.

// Example 1:
//
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
// Example 2:
//
// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
// Example 3:
//
// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
// Example 4:
//
// Input: s = ""
// Output: 0

func lengthOfLongestSubstring(s string) int {
	i, maxLen := 0, 0
	memo := make(map[byte]int)

	for j := 0; j < len(s); j++ {
		memo[s[j]]++

		for memo[s[j]] == 2 && i < j {
			memo[s[i]]--
			i++
		}

		if maxLen < j-i+1 {
			maxLen = j - i + 1
		}
	}

	return maxLen
}
