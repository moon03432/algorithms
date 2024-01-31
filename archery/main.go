package main

import "fmt"

type Key struct {
	m int
	n int
}

var cache = make(map[Key]int)

func main() {
	fmt.Println(possibility(2, 3))   // 4
	fmt.Println(possibility(10, 90)) // 92378
}

func possibility(m, n int) int {
	if n > 10*m {
		return 0
	}

	key := Key{
		m: m,
		n: n,
	}

	if res, ok := cache[key]; ok {
		return res
	}

	if m == 1 {
		if 0 <= n && n <= 10 {
			cache[key] = 1

			return 1
		}

		return 0
	}

	res := 0

	for i := 0; i <= 10; i++ {
		res += possibility(m-1, n-i)
	}

	cache[key] = res

	return res
}
