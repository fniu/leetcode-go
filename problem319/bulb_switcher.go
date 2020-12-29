package problem319

import "math"

func bulbSwitch(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return int(math.Sqrt(float64(n)))

}
