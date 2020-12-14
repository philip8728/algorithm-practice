package main

import (
	"fmt"

	_singleList "github.com/philip8728/algorithm-practice/pkg/linkedlist/singlelist"
)

/*
local import path
mod path github.com/philip8728/algorithm-practice/pkg
*/

func main() {
	list := new(_singleList.ListNode)
	list.AddSingleList(list, 1)
	list.AddSingleList(list, 4)
	list.AddSingleList(list, 3)
	list.AddSingleList(list, 2)
	list.AddSingleList(list, 5)
	list.AddSingleList(list, 2)
	fmt.Println(list)
}
