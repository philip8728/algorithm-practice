package sort

//MergeSort defination
func MergeSort(nums []int, begin, end int) []int {
	if begin >= end {
		return []int{}
	}
	pivot := (begin + end) >> 1
	MergeSort(nums, begin, pivot)
	MergeSort(nums, pivot+1, end)
	merge(nums, begin, pivot, end)
	return nums
}

func merge(nums []int, begin, pivot, end int) {
	temp := make([]int, (end-begin)+1)
	i, j, k := begin, pivot+1, 0
	for i <= pivot && j <= end {
		if nums[i] < nums[j] {
			temp[k] = nums[i]
			k++
			i++
		} else {
			temp[k] = nums[j]
			k++
			j++
		}
	}
	for i <= pivot {
		temp[k] = nums[i]
		k++
		i++
	}
	for j <= end {
		temp[k] = nums[j]
		k++
		j++
	}
	for p := 0; p < len(temp); p++ {
		nums[begin+p] = temp[p]
	}
}
