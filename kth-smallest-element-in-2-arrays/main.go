package main

/*
2个有序数组的第K小元素

两个已经排好序的数组，找出两个数组合并后的第K小的数。
如两个数组[1 2 3 4 5 6] [6 7 8 9 10 11 12] ，K=8
输出：7
*/

import (
	"fmt"
)

func main() {
	var array1 = []int{1, 2, 3, 4, 5, 6}
	var array2 = []int{6, 7, 8, 9, 10, 11, 18}
	var k int = 8

	fmt.Println(kth(array1, array2, k))
}

func kth(a1, a2 []int, k int) int {
	n1, n2 := len(a1), len(a2)
	i, j := 0, 0

	for k > 1 {
		if i >= n1 && j >= n2 {
			return -1
		}

		if i >= n1 {
			j, k = j+1, k-1
			continue
		}

		if j >= n2 {
			i, k = i+1, k-1
			continue
		}

		if a1[i] < a2[j] {
			i, k = i+1, k-1
			continue
		}

		j, k = j+1, k-1
	}

	// k = 1
	if i >= n1 && j >= n2 {
		return -1
	}

	if i >= n1 {
		return a2[j]
	}

	if j >= n2 {
		return a1[i]
	}

	return min(a1[i], a2[j])
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
