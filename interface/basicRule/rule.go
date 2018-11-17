package  main

import (
	"bytes"
	"fmt"
	"os"
)

/*

package io

type Writer interface {
	Write(p []byte) (n int, error)
	}

func Fprintf(w io.Write, format string, args ...interface{}) (int, error)
*/

type byteCounter int

func (c *byteCounter) Write (p []byte) (int, error) {
	*c += byteCounter(len(p))
	return len(p), nil
}


func printf(format string, args ...interface{}) (int, error) {
	return fmt.Fprintf(os.Stdout, format, args)
}

func printfLine(format string, args ...interface{}) string{
	var buf bytes.Buffer
	fmt.Fprintf(&buf, format, args)
	return buf.String()
}


func main() {
	fmt.Println("interface rule")
	var c byteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)
	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello %s", name)
	fmt.Println(c)
	fmt.Println(printf("%s","hello" ))
	fmt.Println(printfLine("%s", "dog"))
}


/*txt

一个类型如果拥有一个接口所需要的所有方法，就可以说这个类型实现了这个接口
也就是这个类型属于这个接口===>所以一个接口可以有很多类型
如果方法接受者是指针类型，要保证变量必须可以寻址，空接口和空结构体无法寻址
空接口类型对实现它的类型没有要求，可以将任意值赋给空接口类型









 */