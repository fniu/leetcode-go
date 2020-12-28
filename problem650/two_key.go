package problem650

func minSteps(n int) int {
	ans := 0
	d := 2

	for n > 1 {
		for n%d == 0 {
			ans += d
			n /= d
		}
		d++
	}

	return ans

}

func minStepsDP(n int) int {
	if n == 1 {
		return 0
	}
	if n <= 3 {
		return n
	}
	var divider float64
	var step, dividerInt int
	minStep := n
	for numPaste := 2; numPaste <= n/2; numPaste++ {
		divider = float64(n) / float64(numPaste)
		dividerInt = int(divider)
		if float64(dividerInt) == divider {
			step = minSteps(dividerInt) + numPaste
			if step < minStep {
				minStep = step
			}
		}
	}
	return minStep
}
