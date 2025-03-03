package main

import "fmt"

func main() {
	//fmt.Println(generateCombination([]string{"a"}))
	//fmt.Println(generateCombination([]string{"a", "b"}))
	//fmt.Println(generateCombination([]string{"a", "b", "c"}))

	fmt.Println(generateCombinationDFS([]string{"a"}))
	fmt.Println(generateCombinationDFS([]string{"a", "b"}))
	fmt.Println(generateCombinationDFS([]string{"a", "b", "c"}))
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

func generateCombinationDFS(a []string) [][]string {
	ch := make(chan []int)

	g := make([]int, 0)

	res := make([][]string, 0)

	go func() {
		dfs(g, len(a), ch)
		close(ch)
	}()

	for {
		data, ok := <-ch
		if !ok {
			return res
		}

		if s := convert(a, data); len(s) > 0 {
			res = append(res, s)
		}

	}
}

func dfs(g []int, n int, ch chan []int) {
	if len(g) == n {
		ch <- g

		return
	}

	dfs(append(g, 0), n, ch)
	dfs(append(g, 1), n, ch)
}

func convert(a []string, g []int) []string {
	res := make([]string, 0)

	for i, exists := range g {
		if exists > 0 {
			res = append(res, a[i])
		}
	}

	return res
}
