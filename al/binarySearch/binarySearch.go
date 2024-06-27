package binarysearch

func BinarySearch(nums []int, target int) (int, bool) {

	beg := 0
	end := len(nums) - 1
	mid := len(nums) / 2

	for beg <= end {

		if nums[mid] == target {
			return mid, true
		} else if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			beg = mid + 1
		}

		mid = (beg + end) / 2
	}

	return -1, false
}

/*
func main() {
	nums := []int{0, 2, 4, 6, 8, 10}
	var target int = 4

	index, ok := BinarySearch(nums, target)

	if !ok {
		fmt.Println("value", target, "not found")
	} else {
		fmt.Println("value found at index", index)
	}
}
*/