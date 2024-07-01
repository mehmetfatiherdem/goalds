package selectionsort

import "testing"

var(
	happyPathArray = []int{64, 25, 12, 22, 11}
	happyPathArrayNested = []int{64, 25, 12, 22, 11}
)

func areArraysEqual(arr1 []int, arr2 []int) (equal bool, ok bool) {
	if len(arr1) != len(arr2) {
		return false, false
	}

	for i, val := range arr1 {
		if arr2[i] != val {
			return false, true
		}
	}

	return true, true
}

func TestHappyPathCase(t *testing.T) {
	sorted := []int{11,12,22,25,64}

	SelectionSort(happyPathArray)

	equal, ok := areArraysEqual(sorted, happyPathArray)

	if !ok {
		t.Fatalf("array sizes are not equal!!!")
	}

	if !equal {
		t.Fatalf("array could not be sorted properly!!")
	}
}


func TestHappyPathCaseNested(t *testing.T) {

	sorted := []int{11,12,22,25,64}

	SelectionSortNestedLoop(happyPathArrayNested)

	equal, ok := areArraysEqual(sorted, happyPathArrayNested)

	if !ok {
		t.Fatalf("array sizes are not equal!!!")
	}

	if !equal {
		t.Fatalf("array could not be sorted properly!!")
	}

}


