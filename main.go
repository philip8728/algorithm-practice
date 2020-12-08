package main

/*
local import path
mod path github.com/philip8728/algorithm-practice/pkg
*/

import (
	"fmt"

	_mergeSort "github.com/philip8728/algorithm-practice/pkg/sort"
)

func main() {
	a := []int{10, 5, 6, 7, 8}
	begin, end := 0, len(a)-1
	_mergeSort.QuickSort(a, begin, end)
	fmt.Println(a)
}
