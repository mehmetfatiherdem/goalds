package insertionSort

func InsertionSort(nums []int) {

	for i := 1; i < len(nums); i++ {
		t := i
		for j := i - 1; j >= 0; j-- {
			if nums[j] > nums[t] {
				nums[j], nums[t] = nums[t], nums[j]
				t = j
			}
		}
	}
}
