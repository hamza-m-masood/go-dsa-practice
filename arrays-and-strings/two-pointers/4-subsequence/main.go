package main

import "fmt"

func main() {
	fmt.Println(checkSubsequence("ace", "abcde"))
}

func checkSubsequence(s string, t string) bool {
	bs := []byte(s)
	ts := []byte(t)

	j := 0
	i := 0
	for i < len(bs) {
		if bs[i] == ts[j] {
			i++
		} else {
			j++
		}
		if j == len(ts) {
			return false
		}
	}
	return true
}
