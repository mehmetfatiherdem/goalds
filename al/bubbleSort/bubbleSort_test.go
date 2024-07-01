package bubblesort

import (
	"testing"

	arrayfuncs "example.com/arrayFuncs"
)

var(
	happyPathArray = []int {6, 0, 3, 5}
)

func TestHappyPathCase(t *testing.T) {
	sorted := []int {0, 3, 5, 6}

	BubbleSort(happyPathArray)

	equal, ok := arrayfuncs.AreArraysEqual(sorted, happyPathArray)

	if !ok {
		t.Fatalf("array sizes are not equal!!!")
	}

	if !equal {
		t.Fatalf("array could not be sorted properly!!")
	}
	
}
