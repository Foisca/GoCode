package main

import (
	"fmt"
)

func main() {
	var i interface{}
	i = 10
	fmt.Printf("type = %T\n", i)
	fmt.Printf("dec:%d\n", i)
	fmt.Printf("bin:%b\n", i)
	fmt.Printf("oct:%O\n", i)
	fmt.Printf("hex:%x\n", i)
}
