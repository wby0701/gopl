package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("commencing countdow")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		j := <-tick
		fmt.Println(j)
	}
	fmt.Println("biu biu biu")
}
