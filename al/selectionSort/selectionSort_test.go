package selectionsort

import "testing"

var(
	happyPathArray = []int{64, 25, 12, 22, 11}
)

func TestHappyPathCase(t *testing.T) {
	sorted := []int{11,12,22,25,64}

	SelectionSort(happyPathArray)

	for i := 0; i<len(sorted); i++ {
		if happyPathArray[i] != sorted[i] {
			t.Fatalf("array could not be sorted properly!!")
		}
	}
}
