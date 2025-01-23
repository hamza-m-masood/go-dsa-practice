package main

import "fmt"

func main() {
	fmt.Println(palindromeCheck("racecar"))
}

func palindromeCheck(input string) bool {
	// converting to arrray of bytes. Each element in the array represents a character in the string.
	byteArray := []byte(input)
	right := len(byteArray) - 1
	left := 0
	for left < right {
		if byteArray[left] != byteArray[right] {
			return false
		}
		left += 1
		right -= 1
	}
	return true
}
