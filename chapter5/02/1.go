package main

import (
	. "fmt"
)

func main() {
	//var a *int
	//Println(a == nil)

	//var aa int
	//bb := & aa
	//Printf("%T\n", bb)
	//cc := & bb
	//Printf("%T\n", cc)

	//var i int
	//p := & i
	//i = 5
	//Println(*p)
	//Println(i)
	//*p = 10
	//Println(*p)
	//Println(i)

	i := 1
	inc(&i)
	inc(&i)
	inc(&i)
	Println(i)
	Printf("%T\n", i)
	Printf("%T\n", &i)

	Println("\n---------\n\n")

	//p := &[3]int{1, 2, 3}
	//Printf("%T\n", p)
	//Printf("%T\n", *p)
	//pow(p)
	//Println(p)

	aa := &[]int{1, 2, 3, 4}
	pow2(aa)
	Println(aa)

}

func inc(p *int) {
	*p++
}

//func pow(p *[3]int) {
//	i := 0
//	for i < len(p) {
//		(*p)[i] = (*p)[i] * (*p)[i]
//		i++
//	}
//}


func pow2(p *[]int) {
	i := 0
	for i < len(*p) {
		(*p)[i] = (*p)[i] * (*p)[i]
		i++
	}
}
