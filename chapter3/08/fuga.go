package main

import (
	"fmt"
)

func main()  {
	//fmt.Println("aaaaaaaaa")

	//aaa := [5]string{
	//	"aa",
	//	"bb",
	//	"cc",
	//	//4,
	//	//5,
	//}
	//fmt.Printf("%v", aaa)

	bbb := [...]int{
		1,
		2,
		3,
	}
	fmt.Printf("%v\n", bbb[2])
	bbb[2] = 100001
	fmt.Printf("%v\n", bbb[2])

	ccc := [...] int8 {
		1,
		2,
		3,
	}
	ccc[0] = 126
	fmt.Printf("%v\n", ccc)

	a1 := [3]int{1,2,3}
	a2 := [3]int{4,5,6}

	a1 = a2

	fmt.Printf("%v", a1)


	//3.9
	var ddd interface{}
	ddd = 1
	ddd = "aaaaaa"
	ddd = [3]int{}
	fmt.Printf("%v", ddd)

	var eee, fff interface{}
	eee = 1
	fff = 2
	fmt.Println(eee , fff)

}