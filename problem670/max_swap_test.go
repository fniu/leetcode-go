package problem670

import (
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {

	var result, expected int

	result = maximumSwap(1234)
	if result != 4231 {
		t.Errorf("Expect %d, got=%d", 4231, result)
	}

	result = maximumSwap(3312)
	expected = 3321
	if result != expected {
		t.Errorf("Expect %d from 3312, got=%d", expected, result)
	}

	result = maximumSwap(3)
	expected = 3
	if result != expected {
		t.Errorf("Expect %d, got=%d", expected, result)
	}
	result = maximumSwap(32)
	expected = 32
	if result != expected {
		t.Errorf("Expect %d, got=%d", expected, result)
	}

	result = maximumSwap(33)
	expected = 33
	if result != expected {
		t.Errorf("Expect %d, got=%d", expected, result)
	}

	result = maximumSwap(9973)
	expected = 9973
	if result != expected {
		t.Errorf("Expect %d from 9973, got=%d", expected, result)
	}

	result = maximumSwap(129973)
	expected = 929173
	if result != expected {
		t.Errorf("Expect %d, got=%d", expected, result)
	}

	result = maximumSwap(98368)
	expected = 98863
	if result != expected {
		t.Errorf("Expect %d, got=%d", expected, result)
	}

	result = maximumSwap(8368)
	expected = 8863
	if result != expected {
		t.Errorf("Expect %d from 8368, got=%d", expected, result)
	}

}

func TestSwap(t *testing.T) {
	var bs []byte
	var num, expected int

	num = 345
	bs = []byte(strconv.Itoa(num))
	expected = 435
	if expected != swapIndexesToInt(bs, 0, 1) {
		t.Errorf("Expected %d, got=%d", expected, swapIndexesToInt(bs, 0, 1))
	}

}

// func TestAddNode(t *testing.T) {
// 	var n, m *SortedList

// 	n = nil
// 	m = addNode(n, byte(10), 0)
// 	fmt.Printf("Got %v\n", m)

// 	m = addNode(m, byte(20), 1)
// 	fmt.Printf("Got %v\n", m)
// 	fmt.Printf("Got %v\n", m.LessEq)

// 	m = addNode(m, byte(15), 2)
// 	fmt.Printf("Got %v\n", m)
// 	fmt.Printf("Got %v\n", m.LessEq)
// 	fmt.Printf("Got %v\n", m.LessEq.LessEq)

// 	m = addNode(m, byte(15), 3)
// 	fmt.Printf("Got %v\n", m)
// 	fmt.Printf("Got %v\n", m.LessEq)
// 	fmt.Printf("Got %v\n", m.LessEq.LessEq)

// }
