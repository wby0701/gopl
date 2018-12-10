package main

import (
	"flag"
	"fmt"
	"gopl/interface/flagValue/flag"
	"time"
)

//var period = flag.Duration("period", 1*time.Second, "sleep period")
var temp = flagmy.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	time.Now()
	flag.Parse()
	fmt.Println(*temp)
	//fmt.Printf("sleeping for %v...", *period)
	//time.Sleep(*period)
	fmt.Println()
}