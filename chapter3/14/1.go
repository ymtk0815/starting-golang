package main

import (
	. "fmt"
)

func main() {

	aaa := 111
	bbb := 222

	if x, y := aaa, bbb; x < y {
		Println("aaaaaaaaaa")
	}

	x, y := 3, 5
	if n := x * y; n%2 == 1 {
		Println("bbbbbbb")
	} else {
		Println("cccccccccc")
	}

	// for
	i := 0
	for {
		//Println("dddd", i)
		i++
		if i == 100 {
			break
		}
	}

	i = 0
	for i < 100 {
		//Println("eeee", i)
		i++
	}

	for j := 0; j < 100; j++ {
		//Println("ffff", j)
	}

	for k := 0; k < 100; k++ {
		if k % 10 == 2 {
			//Println("222222222222", k)
			continue
		}
		//Println("1010101010", k)
	}

	fruits := [3]string{"Apple", "Bananana", "Cherry"}
	for i, s := range fruits {
		//Println("llll", i, s)
	}

}
