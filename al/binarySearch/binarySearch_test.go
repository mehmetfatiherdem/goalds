package binarysearch

import "testing"

var (
	happyPathArray = []int{0, 2, 4, 6, 8, 10}
) 

func TestExistingElementInTheMiddle(t *testing.T) {
	target := 4

	_, ok := BinarySearch(happyPathArray, target)

	if !ok {
		t.Fatalf("value %d is not found in the array", target)
	}
}

func TestFirstElement(t *testing.T) {
	target := 0

	_, ok := BinarySearch(happyPathArray, target)

	if !ok {
		t.Fatalf("value %d is not found in the array", target)
	}
} 

func TestLastElement(t *testing.T) {
	target := 10

	_, ok := BinarySearch(happyPathArray, target)

	if !ok {
		t.Fatalf("value %d is not found in the array", target)
	}
}

func TestNonExistingElement(t *testing.T) {
	target := 3

	_, ok := BinarySearch(happyPathArray, target)

	if ok {
		t.Fatalf("value %d is not in the array but returned ok", target)
	}
}
