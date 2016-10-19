package main

import (
	. "fmt"
)

func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func pow(a [3]int) {
	for i, v := range a {
		a[i] = v * v
	}
	return
}
func pow2(a []int) {
	for i, v := range a {
		a[i] = v * v
	}
	return
}

func main() {
	Println(sum(1, 2, 3, 4, 5))

	aa := []int{1, 2, 3}
	Println(sum(aa...))

	a := [3]int{1, 2, 3}
	pow(a)
	Println(a)

	bb := []int{1, 2, 3}
	pow2(bb)
	Println(bb)

}
