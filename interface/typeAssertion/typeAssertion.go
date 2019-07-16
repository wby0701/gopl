package main

import (
	"bytes"
	"io"
	"os"
)

type base interface {
	test() int
}

type bases interface {
	test() int
	tests() int
}

type Base struct {
	value1 int
}

type Bases struct {
	value2 int
	value3 int
}

func (B *Base) test() int {
	return B.value1
}

func (B *Bases) test() int {
	return B.value2
}

func (B *Bases) tests() int {
	return B.value3
}

func main() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File) // success: f == os.Stdout
	c := w.(*bytes.Buffer) //
	f.Name()
	c.Len()
}