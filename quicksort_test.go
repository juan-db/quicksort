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

func TestQuicksortEmptySlice(t *testing.T) {
	expected := []int{}
	actual := []int{}
	Sort(actual)
	if !equal(expected, actual) {
		t.Errorf("Unexpected result when sorting an empty slice.\nExpected: %v\nGot.....: %v\n", expected, actual)
	}
}

func TestQuicksortSingleElementSlice(t *testing.T) {
	expected := []int{0}
	actual := []int{0}
	Sort(actual)
	if !equal(expected, actual) {
		t.Errorf("Unexpected result when sorting a slice with a single element.\nExpected: %v\nGot.....: %v\n", expected, actual)
	}
}

func TestQuicksortTwoElementSlice(t *testing.T) {
	expected := []int{0, 1}
	actual := []int{0, 1}
	Sort(actual)
	if !equal(expected, actual) {
		t.Errorf("Unexpected result when sorting a slice with two elements that are already sorted.\nExpected: %v\nGot.....: %v\n", expected, actual)
	}
}

func TestQuicksortTwoElementUnsortedSlice(t *testing.T) {
	expected := []int{2, 7}
	actual := []int{7, 2}
	Sort(actual)
	if !equal(expected, actual) {
		t.Errorf("Unexpected result when sorting a slice with two elements that aren't sorted.\nExpected: %v\nGot.....: %v\n", expected, actual)
	}
}

func TestQuicksortSmallSample(t *testing.T) {
	expected := []int{0, 1, 3, 7}
	actual := []int{7, 0, 1, 3}
	Sort(actual)
	if !equal(expected, actual) {
		t.Errorf("Unexpected result when sorting a slice with elements that aren't sorted.\nExpected: %v\nGot.....: %v\n", expected, actual)
	}
}

func TestQuicksort(t *testing.T) {
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7}
	actual := []int{7, 6, 5, 4, 3, 2, 1, 0}
	Sort(actual)
	if !equal(expected, actual) {
		t.Errorf("Unexpected result when sorting a slice with elements that aren't sorted.\nExpected: %v\nGot.....: %v\n", expected, actual)
	}
}

func TestQuicksortSortedSlice(t *testing.T) {
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7}
	actual := []int{0, 1, 2, 3, 4, 5, 6, 7}
	Sort(actual)
	if !equal(expected, actual) {
		t.Errorf("Unexpected result when sorting a slice with elements that are already sorted.\nExpected: %v\nGot.....: %v\n", expected, actual)
	}
}
