package main

import "fmt"

type err struct {
	key 	int
	value 	string
}

func (e err) Error() string {
	fmt.Println(e.value)
	return e.value
}

func main() {
	e := err{
		key : 1,
		value: "qqqq",
	}

	e.Error()
	s := [...]string {
		"aaa",
		"bbb",
		"ccc",
	}
	fmt.Println(len(s))
	fmt.Println(s[1])
}