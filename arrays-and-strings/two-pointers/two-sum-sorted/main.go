package main

import "fmt"

func main() {
	fmt.Println(twoSumSorted([]int{1, 2, 4, 6, 8, 9, 14, 15}, 13))
}

func twoSumSorted(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1
	for left < right {
		curr := nums[left] + nums[right]
		if curr == target {
			return true
		}
		if curr > target {
			right -= 1
		} else {
			left += 1
		}
	}
	return false
}
