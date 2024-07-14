package main

import "fmt"

func merge(nums []int, l int, m int, r int) {
	n1 := m - l + 1
	n2 := r - m

	larr := make([]int, n1)
	rarr := make([]int, n2)

	for i := 0; i < n1; i++ {
		larr[i] = nums[l+i]
	}

	for j := 0; j < n2; j++ {
		rarr[j] = nums[m+1+j]
	}

	k := l
	n := 0
	o := 0

	for n < n1 && o < n2 {
		if larr[n] <= rarr[o] {
			nums[k] = larr[n]
			n++
		} else {
			nums[k] = rarr[o]
			o++
		}

		k++
	}

	for n < n1 {
		nums[k] = larr[n]
		n++
		k++
	}

	for o < n2 {
		nums[k] = rarr[o]
		o++
		k++
	}

}

func MergeSort(nums []int, l int, r int) {
	if l < r {
		m := l + ((r - l) / 2)

		MergeSort(nums, l, m)
		MergeSort(nums, m+1, r)

		merge(nums, l, m, r)
	}
}

func main() {

	arr := []int{38, 27, 43, 3, 9, 82, 10}
	MergeSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted array is:", arr)
}
