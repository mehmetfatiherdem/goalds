package insertionsort

import (
	"testing"

	"goalds/utils"
)

var (
	happyPathArray = []int{23,1,10,5,2}
	sortedArray = []int{1,2,3,4,5}
)

func TestHappyPathCase(t *testing.T) {
	sorted := []int{1,2,5,10,23}

	InsertionSort(happyPathArray)

	isEqual, ok := utils.AreArraysEqual(happyPathArray, sorted)

	if !ok {
		t.Fatalf("arrays are not equal size!!!")
	}

	if !isEqual {
		t.Fatalf("array is not sorted properly!!")
	}
}

func TestSortedArrayCase(t *testing.T) {
	sorted := [] int{1,2,3,4,5}

	InsertionSort(sortedArray)

	isEqual, ok := utils.AreArraysEqual(sortedArray, sorted)

	if !ok {
		t.Fatalf("arrays are not equal size!!!")
	}

	if !isEqual {
		t.Fatalf("array is not sorted properly!!")
	}
}

