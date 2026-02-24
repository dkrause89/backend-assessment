package main

import "sort"

func Solution(A, B, C []int) int {
	n := len(A)
	m := len(C)
	if n == 0 || m == 0 {
		return -1
	}

	low, high := 1, m
	for low < high {
		mid := (low + high) / 2
		if canNailAll(A, B, C, mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	if canNailAll(A, B, C, low) {
		return low
	}
	return -1
}

func canNailAll(A, B, C []int, J int) bool {
	nails := make([]int, J)
	copy(nails, C[:J])
	sort.Ints(nails)

	for k := 0; k < len(A); k++ {
		start, end := A[k], B[k]
		i := sort.Search(len(nails), func(i int) bool { return nails[i] >= start })
		if i >= len(nails) || nails[i] > end {
			return false
		}
	}
	return true
}
