package main

import (
	"fmt"
	"strings"
)

type memo map[int][]int

type combinations [][]int

func bestSum(targetSum int, numbers []int, m memo) ([]int, combinations) {
	if _, exists := m[targetSum]; exists {
		return m[targetSum], nil
	}

	if targetSum == 0 {
		return []int{}, nil
	}

	if targetSum < 0 {
		return nil, nil
	}

	var shortestCombination []int
	var possibleCombinations combinations

	for _, number := range numbers {
		remainder := targetSum - number

		remainderCombination, _ := bestSum(remainder, numbers, m)

		if remainderCombination != nil {
			combination := append(remainderCombination, number)

			if shortestCombination == nil || len(combination) < len(shortestCombination) {
				possibleCombinations = append(possibleCombinations, combination)

				shortestCombination = combination
			}
		}

	}

	m[targetSum] = shortestCombination
	return m[targetSum], possibleCombinations
}

func howSum(targetSum int, numbers []int, m memo) []int {

	if _, exists := m[targetSum]; exists {
		return m[targetSum]
	}

	if targetSum == 0 {
		return []int{}
	}

	if targetSum < 0 {
		return nil
	}

	for _, number := range numbers {
		remainder := targetSum - number

		result := howSum(remainder, numbers, m)

		if result != nil {
			//fmt.Println(targetSum, "-", number, "=", remainder)

			result = append(result, number)
			m[targetSum] = result

			return m[targetSum]
		}
	}

	m[targetSum] = nil
	return m[targetSum]
}

func minExpenses(values []int) int {
	sum := 0

	for _, value := range values {
		fmt.Println(value)
		sum += value
	}

	fmt.Println(sum)

	if sum >= 0 {
		return 0
	}

	return -1
}

func maxValue(s string, column string) int {
	values := strings.Split(s, "\n")

	headers := strings.Split(values[0], ",")

	fmt.Println(headers)

	return 0
}

func main() {
	//fmt.Println(findJudge(3, [][]int{{1,3}, {2,3}})) // 3
	//fmt.Println(findJudge(2, [][]int{{1,2}})) // 2
	//fmt.Println(findJudge(3, [][]int{{1,3}, {2,3}, {3,1}})) // -1
	//fmt.Println(findJudge(1, [][]int{})) // 1
	//fmt.Println(findJudge(2, [][]int{})) // 2

	//fmt.Println(smallestPositiveInteger([]int{1, 3, 6, 4, 1, 2})) // 5
	//fmt.Println(smallestPositiveInteger([]int{1, 2, 3}))          // 4
	//fmt.Println(smallestPositiveInteger([]int{-1, -3}))           // 1
	//fmt.Println(smallestPositiveInteger([]int{1, 2, 0}))          // 1

	//fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [0,1]
	//fmt.Println(twoSum([]int{3, 2, 4}, 6))      // [1,2]
	//fmt.Println(twoSum([]int{3, 3}, 6))         // [0,1]

	//fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
	//fmt.Println(lengthOfLongestSubstring("bbbbb"))    // 1
	//fmt.Println(lengthOfLongestSubstring("pwwkew"))    // 3

	//fmt.Println(bestSum(7, []int{5, 3, 4, 7}, memo{})) // [7] [[4 3] [7]]
	//fmt.Println(bestSum(8, []int{2, 3, 5}, memo{}))    // [5 3] [[3 3 2] [5 3]]
	//fmt.Println(bestSum(8, []int{1, 4, 5}, memo{}))    // [4 4] [[5 1 1 1] [4 4]]
	//fmt.Println(bestSum(100, []int{1, 2, 5, 25}, memo{}))

	//fmt.Println(howSum(7, []int{2, 3}, memo{}))       // [3,2,2]
	//fmt.Println(howSum(7, []int{5, 3, 4, 7}, memo{})) // [4,3]
	//fmt.Println(howSum(7, []int{2, 4}, memo{}))       // []
	//fmt.Println(howSum(8, []int{2, 3, 5}, memo{}))    // [2,2,2,2]
	//fmt.Println(howSum(300, []int{7, 14}, memo{}))    // []

	//fmt.Println(maxValue("id,name,age,act.,room,dep.\n1,Jack,68,T,13,8\n17,Betty,28,F,15,7", "age"))
	//fmt.Println(maxValue("area,land\n3722,CN\n6612,RU\n3855,CA\n3797,USA", "area"))
	//fmt.Println(maxValue("city,temp2,temp\nParis,7,-3\nDubai,4,-4\nPorto,-1,-2", "temp"))
}
