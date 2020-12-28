package problem650

import "testing"

func TestSolution(t *testing.T) {
	step := minSteps(3)
	expectedStep(t, 3, step)
	expectedStep(t, 5, minSteps(5))
	expectedStep(t, 0, minSteps(1))
	expectedStep(t, 4, minSteps(4))
	expectedStep(t, 6, minSteps(8))
	expectedStep(t, 10, minSteps(32))
	expectedStep(t, 11, minSteps(28))
	expectedStep(t, 9, minSteps(27))
}

func expectedStep(t *testing.T, expected, got int) {
	t.Helper()
	if expected != got {
		t.Errorf("Expected %d, got=%d", expected, got)
	}
}
