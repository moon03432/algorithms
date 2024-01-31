package main

import "fmt"

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0)) // 4
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3)) // -1
	fmt.Println(search([]int{1}, 0))                   // -1
}

func search(nums []int, target int) int {
	return search_index(nums, 0, len(nums), target)
}

func search_index(nums []int, i, j, target int) int {
	if j == i {
		return -1
	}

	if nums[i] == target {
		return i
	}

	if j == i+1 {
		return -1
	}

	if j == i+2 {
		if nums[i+1] == target {
			return i + 1
		}

		return -1
	}

	k := (i + j) / 2

	if target == nums[k] {
		return k
	}

	if nums[i] < target && target < nums[k] {
		return search_index(nums, i+1, k, target)
	}

	if nums[i] < target && target < nums[k] {
		return search_index(nums, i+1, k, target)
	}

}
