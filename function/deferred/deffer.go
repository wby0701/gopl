package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	printStack()
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

//延迟函数的调用在释放堆栈之前
func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

/*
func Parse(intput string) (s *Syntax, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("interner error: %v", p)
		} ()
	}
}
*/
