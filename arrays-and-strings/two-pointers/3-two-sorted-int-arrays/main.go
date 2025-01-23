package main

import "fmt"

func main() {
	arr1 := []int{3, 4, 7, 12}
	arr2 := []int{2, 5, 8, 9, 20}
	fmt.Println(combine(arr1, arr2))
}

func combine(arr1 []int, arr2 []int) []int {
	i, j := 0, 0
	ans := make([]int, 0)

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			ans = append(ans, arr1[i])
			i += 1
		} else {
			ans = append(ans, arr2[j])
			j += 1
		}
	}

	// add remaining elements for arr1 if the elments of arr1 are left
	for i < len(arr1) {
		ans = append(ans, arr1[i])
		i++
	}

	// add remaining elements for arr2 if the elments of arr2 are left
	for j < len(arr2) {
		ans = append(ans, arr2[j])
		j++
	}
	return ans
}
