package search

//BinarySearch function
func BinarySearch(nums []int, begin, end, target int) int {
	pivot := (begin + end) >> 1
	if nums[pivot] == target {
		return pivot
	} else if target < nums[pivot] {
		return BinarySearch(nums, 0, pivot-1, target)
	} else if target >= nums[pivot] {
		return BinarySearch(nums, pivot+1, end, target)
	}
	return -1
}
