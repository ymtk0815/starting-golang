package main

import (
	. "fmt"
)

func main() {
	s := make([]int, 10)
	Println(s)

	var a [10]int
	Println(a)

	aaa := make([]float64, 3)
	Println(aaa)
	aaa[0] = 10
	Println(aaa)
	aaa[1] = 6.28
	Println(aaa)

	bbb := make([]int, 8)
	Println(bbb)
	Println(len(bbb))
	ccc := [4]int{1, 2, 3, 4}
	Println(len(ccc))

	Println("\n------------------------------------\n")

	s1 := make([]int, 5)
	Println(s1)
	Println(len(s1))
	Println(cap(s1))

	s2 := make([]int, 5, 10)
	Println(s2)
	Println(len(s2))
	Println(cap(s2))

	Println("\n------------------------------------\n")

	ddd := [5]int{1, 2, 3, 4, 5}
	eee := ddd[0:4]
	Println(eee)


	fff := "ABCDE"[1:3]
	Println(fff)

}
