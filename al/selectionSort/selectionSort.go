package selectionsort

func SelectionSort(nums []int) {
	firstUnsortedIndex := 0
	var lowestIndex int = firstUnsortedIndex

	i := firstUnsortedIndex + 1

	for firstUnsortedIndex < len(nums)-1 {
		if nums[i] < nums[lowestIndex] {
			lowestIndex = i
		}

		if i < len(nums)-1 {
			i = i + 1
			continue
		}

		nums[firstUnsortedIndex], nums[lowestIndex] = nums[lowestIndex], nums[firstUnsortedIndex]

		firstUnsortedIndex++
		lowestIndex = firstUnsortedIndex
		i = firstUnsortedIndex + 1

	}
}
