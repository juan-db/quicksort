package quicksort

import (
	"math/rand"
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
	actual := Sort([]int{})
	if !equal(expected, actual) {
		t.Errorf("Unexpected result when sorting an empty slice.\nExpected: %v\nGot.....: %v\n", expected, actual)
	}
}

func TestQuicksortSingleElementSlice(t *testing.T) {
	expected := []int{0}
	actual := Sort([]int{0})
	if !equal(expected, actual) {
		t.Errorf("Unexpected result when sorting a slice with a single element.\nExpected: %v\nGot.....: %v\n", expected, actual)
	}
}

func TestQuicksortTwoElementSlice(t *testing.T) {
	expected := []int{0, 1}
	actual := Sort([]int{0, 1})
	if !equal(expected, actual) {
		t.Errorf("Unexpected result when sorting a slice with two elements that are already sorted.\nExpected: %v\nGot.....: %v\n", expected, actual)
	}
}

func TestQuicksortTwoElementUnsortedSlice(t *testing.T) {
	expected := []int{2, 7}
	actual := Sort([]int{7, 2})
	if !equal(expected, actual) {
		t.Errorf("Unexpected result when sorting a slice with two elements that aren't sorted.\nExpected: %v\nGot.....: %v\n", expected, actual)
	}
}

func TestQuicksortSmallSample(t *testing.T) {
	expected := []int{0, 1, 3, 7}
	actual := Sort([]int{7, 0, 1, 3})
	if !equal(expected, actual) {
		t.Errorf("Unexpected result when sorting a slice with elements that aren't sorted.\nExpected: %v\nGot.....: %v\n", expected, actual)
	}
}

func TestQuicksort(t *testing.T) {
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7}
	actual := Sort([]int{7, 6, 5, 4, 3, 2, 1, 0})
	if !equal(expected, actual) {
		t.Errorf("Unexpected result when sorting a slice with elements that aren't sorted.\nExpected: %v\nGot.....: %v\n", expected, actual)
	}
}

func TestQuicksortSortedSlice(t *testing.T) {
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7}
	actual := Sort([]int{0, 1, 2, 3, 4, 5, 6, 7})
	if !equal(expected, actual) {
		t.Errorf("Unexpected result when sorting a slice with elements that are already sorted.\nExpected: %v\nGot.....: %v\n", expected, actual)
	}
}

func BenchmarkSort(b *testing.B) {
	random := rand.New(rand.NewSource(777))
	actual := make([]int, 10000)
	for i := range actual {
		actual[i] = random.Int()
	}
	b.ResetTimer()
}
