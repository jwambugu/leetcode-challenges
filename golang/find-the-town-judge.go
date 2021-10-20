package main

import "fmt"

// In a town, there are n people labeled from 1 to n. There is a rumor that one of these people is secretly the town judge.

// If the town judge exists, then:
// 1. The town judge trusts nobody.
// 2. Everybody (except for the town judge) trusts the town judge.
// 3. There is exactly one person that satisfies properties 1 and 2.

// You are given an array trust where trust[i] = [a, b] representing that the person labeled a trusts the person labeled b.

// Return the label of the town judge if the town judge exists and can be identified, or return -1 otherwise.

// Input: n = 2, trust = [[1,2]]
// Output: 2

// Input: n = 3, trust = [[1,3],[2,3]]
// Output: 3

// Input: n = 3, trust = [[1,3],[2,3],[3,1]]
// Output: -1

func findJudge(n int, trust [][]int) int {
	trustCounts := make(map[int]int)

	fmt.Println(trustCounts)

	for i := 1; i <= n; i++ {
		trustCounts[i] = 0
	}

	fmt.Println(trustCounts, trust)

	for i := 0; i < len(trust); i++ {
		fmt.Println(trust[i][0], trust[i][1])
		trustCounts[trust[i][0]] -= 1
		trustCounts[trust[i][1]] += 1
	}

	for judge, count := range trustCounts {
		if count == n-1 {
			return judge
		}
	}

	if len(trust) == 0 {
		return n
	}

	return -1
}
