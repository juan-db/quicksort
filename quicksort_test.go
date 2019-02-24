package quicksort

import (
	"math/rand"
	"sort"
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
	data := []int{7, 2}
	actual := Sort(data)

	if equal(data, actual) {
		t.Errorf("Array was sorted in place instead of creating a new slice.")
	}

	if !equal(expected, actual) {
		t.Errorf("Unexpected result when sorting a slice with two elements that aren't sorted.\nExpected: %v\nGot.....: %v\n", expected, actual)
	}
}

func TestQuicksortSmallSample(t *testing.T) {
	expected := []int{0, 1, 3, 7}
	data := []int{7, 0, 1, 3}
	actual := Sort(data)

	if equal(data, actual) {
		t.Errorf("Array was sorted in place instead of creating a new slice.")
	}

	if !equal(expected, actual) {
		t.Errorf("Unexpected result when sorting a slice with elements that aren't sorted.\nExpected: %v\nGot.....: %v\n", expected, actual)
	}
}

func TestQuicksort(t *testing.T) {
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7}
	data := []int{7, 6, 5, 4, 3, 2, 1, 0}
	actual := Sort(data)

	if equal(data, actual) {
		t.Errorf("Array was sorted in place instead of creating a new slice.")
	}

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

func testRandomSample(t *testing.T, size int) {
	random := rand.New(rand.NewSource(777))
	data := make([]int, size)
	for i := range data {
		data[i] = random.Int()
	}
	expected := append([]int{}, data...)
	sort.Ints(expected)
	actual := Sort(data)
	if !equal(actual, expected) {
		t.Errorf("Unexpected result when sorting slice with %v random values.", size)
	}
}

func TestQuicksortTinyRandomSample(t *testing.T) {
	testRandomSample(t, 10)
}

func TestQuicksortSmallRandomSample(t *testing.T) {
	testRandomSample(t, 100)
}

func TestQuicksortRandomSample(t *testing.T) {
	testRandomSample(t, 1000)
}

func TestQuicksortMediumRandomSample(t *testing.T) {
	testRandomSample(t, 10000)
}

func TestQuicksortLargeRandomSample(t *testing.T) {
	testRandomSample(t, 100000)
}

func TestQuicksortMassiveRandomSample(t *testing.T) {
	testRandomSample(t, 1000000)
}

func BenchmarkSort(b *testing.B) {
	random := rand.New(rand.NewSource(777))
	actual := make([]int, 10000)
	for i := range actual {
		actual[i] = random.Int()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sort(actual)
	}
}
