package main

import (
	. "fmt"
)

func main() {
	Println("aaa")

	aa := map[int]string {
		1: "Apple",
		2: "Banana",
		3: "Cherry",
	}
	for k, v := range aa {
		Println(k, v)
	}

	Println("\n------------------------------\n\n\n\n\n\n")
	bb := map[string]int {
		"Apple" : 1,
		"Banana" : 2,
		"Cherry" : 3,
	}
	bb["Dfafa"] = 4
	bb["Effefafa"] = 5
	for k, v := range bb {
		Println(k, v)
	}
	Println(len(bb))
	delete(bb, "Apple")
	Println(bb)
	Println(len(bb))

}
