package main

import "fmt"

func main() {
	fmt.Println(generateCombination([]string{"a"}))
	fmt.Println(generateCombination([]string{"a", "b"}))
	fmt.Println(generateCombination([]string{"a", "b", "c"}))
}

func generateCombination(a []string) [][]string {
	res := make([][]string, 0)

	if len(a) == 0 {
		return res
	}

	if len(a) == 1 {
		res = append(res, a)

		return res
	}

	last := generateCombination(a[:len(a)-1])
	for _, e := range last {
		c := make([]string, len(e))
		copy(c, e)
		res = append(res, e, append(c, a[len(a)-1]))
	}

	res = append(res, []string{a[len(a)-1]})

	return res
}
