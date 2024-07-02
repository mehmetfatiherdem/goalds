package insertionSort

import (
	"testing"

	arrayfuncs "example.com/arrayFuncs"
)

var (
	happyPathArray = []int{23,1,10,5,2}
)

func TestHappyPathCase(t *testing.T) {
	sorted := []int{1,2,5,10,23}

	InsertionSort(happyPathArray)

	isEqual, ok := arrayfuncs.AreArraysEqual(happyPathArray, sorted)

	if !ok {
		t.Fatalf("arrays are not equal size!!!")
	}

	if !isEqual {
		t.Fatalf("array is not sorted properly!!")
	}
}

