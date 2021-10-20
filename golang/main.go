package main

import "fmt"

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

	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [0,1]
	fmt.Println(twoSum([]int{3, 2, 4}, 6))      // [1,2]
	fmt.Println(twoSum([]int{3, 3}, 6))         // [0,1]
}
