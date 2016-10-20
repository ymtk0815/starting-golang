package main

import (
	. "fmt"
)

func main() {
	is := [3]int{1, 2, 3}
	ip := & is[1]
	Println(*ip)
	*ip = 0
	Println(*ip)
	Println(is)


	fs := []float64{1.1, 2.2, 3.3}
	fp := & fs[2]
	*fp = 3.14
	Println(fs)

}