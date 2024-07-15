package bubblesort

import (
	"testing"

	"goalds/utils"
)

var(
	happyPathArray = []int {6, 0, 3, 5}
)

func TestHappyPathCase(t *testing.T) {
	sorted := []int {0, 3, 5, 6}

	BubbleSort(happyPathArray)

	equal, ok := utils.AreArraysEqual(sorted, happyPathArray)

	if !ok {
		t.Fatalf("array sizes are not equal!!!")
	}

	if !equal {
		t.Fatalf("array could not be sorted properly!!")
	}
	
}
