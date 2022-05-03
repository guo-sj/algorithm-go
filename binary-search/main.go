// Example code for https://guo-sj.github.io/algorithm/2022/04/16/binary-search.html
package main

import "fmt"

func main() {
	a := []int{1, 2, 4, 9, 30}
	fmt.Println(binarySearch(a, 30, 0))
	fmt.Println(binarySearchLoop(a, 30))
}

// binary search recursive version
func binarySearch(list []int, target int, offset int) int {
	low, high := 0, len(list)-1
	if low > high {
		return -1
	}

	mid := (low + high) / 2
	if list[mid] == target {
		return mid + offset
	}
	if list[mid] > target {
		return binarySearch(list[:mid], target, 0)
	} else {
		return binarySearch(list[mid+1:], target, offset+mid+1)
	}
}

// binary search loop version
func binarySearchLoop(list []int, target int) int {
	low, high := 0, len(list)-1
	for low <= high {
		mid := (low + high) / 2
		if list[mid] == target {
			return mid
		}
		if list[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
