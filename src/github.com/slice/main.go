package main

import "fmt"

func main() {
	var c [3]int
	fmt.Println(c)
	a := []int{1, 2, 3, 4, 5, 6}
	b := append(a[:2], a[3:]...)
	fmt.Println(b)
	fmt.Println(a)
}
