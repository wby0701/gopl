package main

import "fmt"

type IntList struct {
	Value int
	Tail *IntList
}

func (list *IntList) sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.sum()
}

func (list *IntList) create(value int) {
	if list.Tail == nil {
		list.Tail = new(IntList)
		list.Tail.Value = value
	} else {
		list.Tail.create(value)
	}
}

func main() {
	root := IntList{100, nil}
	//var value [6]int = [6]int{1,3, 5, 6, 7, 8}
	q := [...]int{1, 9, 3, 6, 8}
	for _ ,v := range q {
		root.create(v)
	}
	fmt.Println(root)
	fmt.Println(root.sum())
}