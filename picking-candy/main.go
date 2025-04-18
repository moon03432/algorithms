package main

import "fmt"

var (
	indexesOfLength = make(map[int][]int)
)

func main() {
	candies := []int{2, 5, 1, 1, 9, 100, 99}

	indexes := pick(candies)

	fmt.Println(indexes)
	fmt.Println(getSum(candies, indexes))
}

func pick(candies []int) (indexes []int) {
	if res, ok := indexesOfLength[len(candies)]; ok {
		return res
	}

	if len(candies) == 0 {
		return indexes
	}

	if len(candies) == 1 {
		indexesOfLength[1] = []int{0}

		return []int{0}
	}

	if len(candies) == 2 {
		if candies[0] > candies[1] {
			indexesOfLength[1] = []int{0}

			return []int{0}
		}

		indexesOfLength[1] = []int{1}

		return []int{1}
	}

	res1 := append(pick(candies[0:len(candies)-2]), len(candies)-1)
	res2 := pick(candies[0 : len(candies)-1])

	if getSum(candies, res1) > getSum(candies, res2) {
		indexesOfLength[len(candies)] = res1

		return res1
	}

	indexesOfLength[len(candies)] = res2

	return res2
}

func getSum(candies []int, indexes []int) (sum int) {
	for _, i := range indexes {
		sum += candies[i]
	}

	return sum
}
