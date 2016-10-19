package main

import (
	. "fmt"
)

func main() {
	aaa := []int{1, 2, 3}
	bbb := []int{10, 20, 30}
	for i := 0; i < len(bbb); i++ {
		aaa = append(aaa, bbb[i])
	}
	//Println(aaa)

	ccc := []int{11, 22, 33}
	ddd := []int{101, 202, 303}
	ccc = append(ccc, ddd...)
	Println(ccc)
	Println("\n----------------- \n")

	aa := make([]int, 0, 0)
	Println(aa)
	Printf("(A)aa len= %d, cap= %d\n", len(aa), cap(aa))
	aa = append(aa, 1)
	Println(aa)
	Printf("(B)aa len= %d, cap= %d\n", len(aa), cap(aa))
	bb := []int{1, 2, 3, 4, 5}
	aa = append(aa, bb...)
	Println(aa)
	Printf("(C)aa len= %d, cap= %d\n", len(aa), cap(aa))
	aa = append(aa, 1)
	Println(aa)
	Printf("(D)aa len= %d, cap= %d\n", len(aa), cap(aa))
	aa = append(aa, 1)
	Println(aa)
	Printf("(E)aa len= %d, cap= %d\n", len(aa), cap(aa))
	aa = append(aa, 1)
	Println(aa)
	Printf("(F)aa len= %d, cap= %d\n", len(aa), cap(aa))

	Println("\nCopy----------------- \n")

	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{10, 11, 12, 13, 14, 15, 16}
	s3 := copy(s1, s2)
	Println(s1)
	Println(s2)
	Println(s3)

	Println("\n完全----------------- \n")

	cc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	dd := cc[2:4]
	Println(dd)
	Println(len(dd))
	Println(cap(dd))

	ee := cc[2:4:4]
	Println(ee)
	Println(len(ee))
	Println(cap(ee))

	Println("\for----------------- \n")
	ff := []string{"Apple", "Banana", "Cherry"}
	for i, v := range ff {
		Println(i, v)
	}


}
