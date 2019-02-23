package quicksort

import (
	"testing"
)

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestQuicksort(t *testing.T) {
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7}
	actual := Sort([]int{7, 6, 5, 4, 3, 2, 1, 0})
	if !equal(expected, actual) {
		t.Errorf("Array was not sorted correctly.\nExpected: %v\nGot.....: %v\n", expected, actual)
	}
}