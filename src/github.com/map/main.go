package main

import "fmt"

func main() {
	statePopulations := make(map[string]int, 10) //data type, length
	statePopulations = map[string]int{
		"California": 1,
		"Texas":      2,
		"New York":   3,
	}
	// the order is not sure
	// slice cannot be the key to a map!
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Texas"])
	statePopulations["Dalian"] = 100
	fmt.Println(statePopulations["Dalian"])

}
