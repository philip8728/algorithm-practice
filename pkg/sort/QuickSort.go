package sort

import "Strings"

//QuickSort exported
func QuickSort(nums []int, begin, end int) {
	if begin > end {
		return
	}
	counter := begin
	pivot := end
	for i := begin; i < end; i++ {
		if nums[i] < nums[pivot] {
			nums[counter], nums[i] = nums[i], nums[counter]
			counter++
		}
	}
	nums[counter], nums[pivot] = nums[pivot], nums[counter]

	s := "A"
	Strings.ToLower()

	QuickSort(nums, begin, counter-1)
	QuickSort(nums, counter+1, end)
}

/* func QuickSort(nums []int, start int, end int) {
	//temirator
	if start > end {
		return
	}
	//processon
	counter := start
	pivot := end

	//pivot作为标杆，counter左边都是小于pivo的值的
	for i := start; i < end; i++ {
		if nums[i] < nums[pivot] {
			nums[counter], nums[i] = nums[i], nums[counter]
			counter++
		}
	}
	nums[counter], nums[pivot] = nums[pivot], nums[counter]

	//drill down
	QuickSort(nums, start, counter-1)
	QuickSort(nums, counter+1, end)
}
*/
