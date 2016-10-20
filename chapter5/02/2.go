package main

import (
	. "fmt"
)

func main() {

	zz := &[3]int{1, 2, 3}
	Println((*zz)[0])
	Println(zz[0])


	a := [3]string{"Apple", "Banana", "Cherry"}
	p := &a
	Println(a[1])
	Println((*p)[1])

	p[2] = "Grape"
	Println(a[2])
	Println(p[2])

	aa := &[3]int{1, 2, 3}
	Println(len(aa))
	Println(cap(aa))
	Println(aa[0:2])

	bb := &[3]string{"Apple", "Banana", "Cherry"}
	for i, v := range bb {
		Println(i, v)
	}
	Printf("%p", bb)

}

