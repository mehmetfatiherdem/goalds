package quicksort

func partition(nums []int, low int, high int) int {

	pi := nums[high]

	i := low - 1

	for j := low; j <= high-1; j++ {
		if nums[j] < pi {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	i++

	nums[i], nums[high] = nums[high], nums[i]
	return i
}

func QuickSort(nums []int, low int, high int) {

	if low < high {
		pi := partition(nums, low, high)

		QuickSort(nums, low, pi-1)
		QuickSort(nums, pi+1, high)
	}
}


