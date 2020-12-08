package main

import (
	"fmt"

	_quickSort "github.com/philip8728/algorithm-practice/pkg/sort"
)

func main() {
	a := []int{5, 3, 1, 7, 8}
	_quickSort.QuickSort(a, 0, len(a)-1)
	fmt.Println(a)
}
