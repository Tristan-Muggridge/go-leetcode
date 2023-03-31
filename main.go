package main

import "fmt"

func twoSum(整数達 []int, ターゲット int) []int {
	辞書 := make(map[int]int)

	for i := 0; i < len(整数達); i++ {
		if value, ok := 辞書[ターゲット-整数達[i]]; ok {
			return []int{value, i}
		} else {
			辞書[整数達[i]] = i
		}
	}
	return []int{0, 0}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
