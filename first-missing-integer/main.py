# 有一个整数数组 8，4，2，5，6，3，9，12，11 求：第一个间断的数 7 时间复杂度o(N) 空间复杂度尽可能小。

def solve(nums: list) -> int:
	m,s = min(nums),set(nums)
	
	i = m
	while i in s:
		i += 1

	return i

print(solve([8,4,2,5,6,3,9,12,11]))