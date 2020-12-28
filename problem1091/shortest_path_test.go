package problem1091

import (
	"testing"
)

func TestSolution(t *testing.T) {

	assertEqual(t, -1, shortestPathBinaryMatrix(
		[][]int{{0, 1, 0}, {1, 1, 0}, {1, 1, 0}}))

	assertEqual(t, 2, shortestPathBinaryMatrix(
		[][]int{{0, 1}, {1, 0}}))

	assertEqual(t, -1, shortestPathBinaryMatrix(
		[][]int{{0, 0, 0, 0, 0}, {1, 1, 0, 1, 0}, {0, 1, 1, 1, 1}, {1, 1, 1, 1, 0}, {0, 1, 1, 0, 0}}))

	assertEqual(t, 5, shortestPathBinaryMatrix(
		[][]int{{0, 1, 1, 0}, {1, 0, 0, 0}, {1, 1, 1, 0}, {1, 1, 1, 0}}))

	assertEqual(t, 2, shortestPathBinaryMatrix([][]int{{0, 1}, {1, 0}}))

	assertEqual(t, 4, shortestPathBinaryMatrix(
		[][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}))

	assertEqual(t, 6, shortestPathBinaryMatrix(
		[][]int{{0, 1, 1, 0}, {1, 0, 1, 1}, {0, 1, 1, 0}, {1, 0, 0, 0}}))

	assertEqual(t, 1, shortestPathBinaryMatrix(
		[][]int{{0}}))

	assertEqual(t, -1, shortestPathBinaryMatrix(
		[][]int{{0, 0, 0, 0, 1}, {0, 1, 0, 1, 0}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 0}}))

	assertEqual(t, 4, shortestPathBinaryMatrix(
		[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))

}

func assertEqual(t *testing.T, expected, got int) {
	t.Helper()
	if expected != got {
		t.Errorf("Expected %d, got=%d", expected, got)
	}
}
