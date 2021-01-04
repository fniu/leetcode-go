package problem670

import (
	"math"
	"strconv"
)

// SortedList binary tree
type SortedList struct {
	Val    byte
	Index  int
	LessEq *SortedList
}

func maximumSwap(num int) int {
	if num <= 11 {
		return num
	}
	if num < 100 {
		d2 := num / 10
		d1 := num % 10
		if d1 <= d2 {
			return num
		}
		return d1*10 + d2
	}

	numBytes := []byte(strconv.Itoa(num))

	numLen := len(numBytes)
	root := &SortedList{Val: numBytes[numLen-1], Index: numLen - 1}

	for i := numLen - 2; i >= 0; i-- {
		root = addNode(root, numBytes[i], i)
	}
	p := root
	for i, v := range numBytes {
		if p.Val > v {
			return swapIndexesToInt(numBytes, i, p.Index)
		}
		if p.Val == v {
			if p.LessEq != nil && p.LessEq.Val == p.Val {
				p.LessEq.Index = p.Index
			}
			p = p.LessEq
		}
	}

	return num
}

func swapIndexesToInt(numBytes []byte, a, b int) int {
	numBytes[a], numBytes[b] = numBytes[b], numBytes[a]
	return bytesToInt(numBytes)
}

func bytesToInt(numBytes []byte) int {
	var re int
	lenBytes := len(numBytes)
	for i, v := range numBytes {
		re += (int(v) - 48) * int(math.Pow10(lenBytes-i-1))
	}
	return re
}
func addNode(tree *SortedList, val byte, index int) *SortedList {
	if tree == nil {
		return &SortedList{Val: val, Index: index}
	}
	if val > tree.Val || (val == tree.Val && index > tree.Index) {
		return &SortedList{Val: val, Index: index, LessEq: tree}
	}
	if val < tree.Val || (val == tree.Val && index < tree.Index) {
		tree.LessEq = addNode(tree.LessEq, val, index)
	}
	return tree
}

func maximumSwapSample(num int) int {
	if num <= 11 {
		return num
	}

	digits := make([]byte, 0, 9)

	for num > 0 {
		digits = append(digits, byte(num%10))
		num /= 10
	}

	for i := len(digits) - 1; i > 0; i-- {
		maxIndex := 0
		for j := 1; j < i; j++ {
			if digits[j] > digits[maxIndex] {
				maxIndex = j
			}
		}

		if digits[maxIndex] > digits[i] {
			digits[i], digits[maxIndex] = digits[maxIndex], digits[i]

			break
		}
	}

	result := int(digits[len(digits)-1])

	for i := len(digits) - 2; i >= 0; i-- {
		result *= 10
		result += int(digits[i])
	}

	return result
}
