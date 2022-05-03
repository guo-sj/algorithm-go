// Example code for https://guo-sj.github.io/algorithm/2022/04/16/quick-sort.html
package main

import "fmt"

func main() {
	a := []int{9, 7, 3, 6, 8, 7}
	quickSort(a)
	fmt.Println(a)
}

func quickSort(list []int) {
	if len(list) < 2 {
		return
	}
	i, j := 0, len(list)-1
	pivot := list[i]
	for i < j {
		for i < j && list[j] > pivot {
			j--
		}
		if i < j {
			list[i] = list[j]
			i++
		}
		for i < j && list[i] <= pivot {
			i++
		}
		if i < j {
			list[j] = list[i]
			j--
		}
	}
	list[i] = pivot
	quickSort(list[:i])
	quickSort(list[i+1:])
}
