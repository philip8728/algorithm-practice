package main

import (
	"fmt"

	_mergeSort "github.com/philip8728/algorithm-practice/pkg/sort"
)

/*
local import path
mod path github.com/philip8728/algorithm-practice
*/

func main() {
	a := []int{10, 5, 6, 7, 8}
	begin, end := 0, len(a)-1
	fmt.Println(_mergeSort.MergeSort(a, begin, end))
}
