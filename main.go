package main

/*
local import path
mod path github.com/philip8728/algorithm-practice/pkg
*/

import (
	"fmt"

	_search "github.com/philip8728/algorithm-practice/pkg/search"
)

func main() {
	a := []int{10, 5, 6, 7, 8}
	target := 6
	fmt.Println(_search.BinarySearch(a, 0, len(a)-1, target))
}
