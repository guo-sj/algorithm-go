package main

import "fmt"

func main() {
	fmt.Println(longestSubstring("clues", "blue"))
	fmt.Println(longestSubstring("hish", "vista"))
	fmt.Println(longestCommonSubsequence("fosh", "fish"))
}

// longestSubstring returns longest substring in s and t
func longestSubstring(s, t string) (string, int) {
	var length int
	var lastValue, maxValue, maxValueIndex int

	if len(s) > len(t) {
		length = len(t)
	} else {
		length = len(s)
	}
	arr := make([]int, length, length)
	for i := 0; i < length; i++ {
		if s[i] == t[i] {
			arr[i] = lastValue + 1
			if arr[i] > maxValue {
				maxValue = arr[i]
				maxValueIndex = i
			}
		} else {
			arr[i] = 0
		}
		lastValue = arr[i]
	}
	if maxValue > 0 {
		return s[maxValueIndex-maxValue+1 : maxValueIndex+1], maxValue
	}
	return "", 0
}

// longestCommonSubsequence returns the longest common subsequence in s and t
func longestCommonSubsequence(s, t string) (string, int) {
	commonSubsequence := []byte{}
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			if t[j] == s[i] {
				commonSubsequence = append(commonSubsequence, s[i])
				break
			}
		}
	}
	return string(commonSubsequence), len(commonSubsequence)
}
