package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

//print interface type
func main() {
	var w io.Writer
	fmt.Printf("%T\n", w)
	w = os.Stdout
	fmt.Printf("%T\n", w)
	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w)
	w = io.Writer(os.Stderr)
	fmt.Printf("%T\n", w)
	w = nil
	fmt.Printf("%T\n", w)
}