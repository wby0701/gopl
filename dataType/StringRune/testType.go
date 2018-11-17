package main

import (
	"fmt"
	"reflect"
)

func printValue() {
	str := "hello, 世界"

	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	fmt.Println("====================")
	for _, ch := range str {
		fmt.Println(ch)
	}
}

func printType() {
	str := "hello, 世界"
	//utf-8 遍历
	for i := 0; i < len(str); i++ {
		ch := str[i]
		chType := reflect.TypeOf(ch)
		fmt.Printf("%s", chType)
	}
	fmt.Println("====================")
	//unicode 遍历
	for _, ch1 := range str {
		ch1Type := reflect.TypeOf(ch1)
		fmt.Printf("%s", ch1Type)
	}
}

func printString() {
	str := "hello, 世界"

	for _, ch := range str {
		fmt.Printf("%q", ch) //"%q"表示''字符字面值
		fmt.Println(string(ch))
	}
}

func main() {
	printValue()
	printType()
	printString()
}
