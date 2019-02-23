package quicksort

import "sync"

func Sort(toSort []int) []int {
	if toSort == nil {
		return nil
	} else if len(toSort) < 2 {
		return toSort
	}

	output := append([]int{}, toSort...)
	sortInPlace(output)
	return output
}

func sortInPlace(toSort []int) {
	if length := len(toSort); length < 2 {
		return
	} else if length < 3 {
		if toSort[0] > toSort[1] {
			toSort[1], toSort[0] = toSort[0], toSort[1]
			return
		}
	}

	compare, pivot := 0, len(toSort) - 1
	for compare < pivot {
		if toSort[compare] > toSort[pivot] {
			toSort[compare], toSort[pivot - 1], toSort[pivot] = toSort[pivot - 1], toSort[pivot], toSort[compare]
			pivot -= 1
		} else {
			compare += 1
		}
	}

	wait := sync.WaitGroup{}
	if compare > 1 {
		go func() {
			sortInPlace(toSort[:compare])
			wait.Done()
		}()
		wait.Add(1)
	}
	if pivot < len(toSort) - 2 {
		go func() {
			sortInPlace(toSort[pivot + 1:])
			wait.Done()
		}()
		wait.Add(1)
	}
	wait.Wait()
}

// Bad explanation and writing follows (it was written mainly for myself to help me understand what I'm thinking):
// At the start we select our pivot (in our case this will be the last element in the array).
// In the array [ 3, 7, 1, 0 ] our pivot will be set to index 3 where the value 0 resides.
// At this point we know that none of the value ahead of the pivot have been sorted. Our first step is to compare the
// first element in the array (index 0, value 3) with the pivot.
// We can see that the value at our compare point (3) is larger than the value at our pivot (0). This means that 3 must
// come after 0 in the array.
// As we mentioned previously, we know that none of the values in the array have been sorted yet. Because of this it
// doesn't matter how we swap values around, as long as the 0 and 3 swap places.
// Once values start getting sorted, we must make sure to keep sorted values in their sorted state.
// To explain this, let's run through the example:
// [ 3, 7, 1, 0 ] - pivot: (index 3, value 0) - compare (index 0, value 3)
// Move our pivot one place forward in the array to make space for the 3 and end up with the following:
// [ 3, 7, 0, 1 ]
// We can swap the pivot with the value before it since the value before it has not been sorted yet. Now let's move the
// 3 after our pivot:
// [ 1, 7, 0, 3 ] - pivot: (index 2, value 0) - compare (index 0, value 1)
// Since we put a new value at index 0, our compare doesn't move. Our pivot does move however.
// [ 1, 0, 7, 3 ]
// [ 7, 0, 1, 3 ] - pivot: (index 1, value 0) - compare (index 0, value 7)
// [ 0, 7, 1, 3 ]
// Our pivot is now in its final location, so we split the array into 3, values before the pivot, the pivot, and values
// after the pivot: [], [ 0 ], [ 7, 1, 3 ]. Now our algorithm becomes Sort([]) + Sort([ 0 ]), + Sort([ 7 1 3 ])
// Since the first and second arrays are sorted, I'm not gonna walk through them, so let's take a look at the third
// array.
// [ 7, 1, 3 ] - pivot: (index 2, value 3) - compare (index 0, value 7)
// [ 7, 3, 1 ]
// [ 1, 3, 7 ] - pivot: (index 1, value 3) - compare (index 0, value 1)
// Sorted.


// [ 7, 3, 4, 1, 0, 2 ] - pivot (index 5, value 2) - compare (index 0, value 7)
// [ 7, 3, 4, 1, 2, 0 ]
// [ 0, 3, 4, 1, 2, 7 ] - pivot (index 4. value 2) - compare (index 0, value 0)
// [ 0, 3, 4, 1, 2, 7 ] - pivot (index 4, value 2) - compare (index 1, value 3)
// [ 0, 3, 4, 2, 1, 7 ]
// [ 0, 1, 4, 2, 3, 7 ] - pivot (index 3, value 2) - compare (index 1, value 1)
// [ 0, 1, 4, 2, 3, 7 ] - pivot (index 3, value 2) - compare (index 2, value 4)
// [ 0, 1, 2, 4, 3, 7 ]
// [ 0, 1 ] + [ 2 ] + [ 4, 3, 7 ]
