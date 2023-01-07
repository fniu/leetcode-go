package problem77

func combine(n int, k int) [][]int {
	if n == k {
		a := []int{}
		for i := 1; i <= n; i++ {
			a = append(a, i)
		}
		return append([][]int{}, a)
	}
	result := [][]int{}
	if k == 1 {
		for i := 1; i <= n; i++ {
			result = append(result, []int{i})
		}
		return result
	}
	stack := [][]int{}
	for i := 1; i < n; i++ {
		stack = append(stack, []int{i})
	}
	for len(stack) > 0 {
		pop := stack[len(stack)-1]
		m := n - k + len(pop) + 1
		stack = stack[:len(stack)-1]
		c := 1
		next := 0
		for {
			next = pop[len(pop)-1] + c
			if next <= m {
				pop2 := make([]int, len(pop))
				copy(pop2, pop)
				pop2 = append(pop2, next)
				if len(pop2) < k {
					stack = append(stack, pop2)
				} else {
					result = append(result, pop2)
				}
			} else {
				break
			}
			c++
		}

	}
	return result

}
