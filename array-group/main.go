package main

import "fmt"

var array = []int{9, 7, 0, -4, 5, -2}

func main() {
	group()

	fmt.Println(array)
}

func group() {
	if len(array) <= 1 {
		return
	}

	i, j := 0, len(array)-1

	for i < j {
		i, j = advance(i, j)
		swap(i, j)
	}
}

func advance(i, j int) (int, int) {
	for i < j && array[i] < 0 {
		i++
	}

	for j > i && array[j] > 0 {
		j--
	}

	return i, j
}

func swap(i, j int) {
	array[i], array[j] = array[j], array[i]
}
