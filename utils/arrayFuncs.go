package arrayfuncs

func AreArraysEqual(arr1 []int, arr2 []int) (equal bool, ok bool) {
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
