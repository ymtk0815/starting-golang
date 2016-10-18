package main

import (
	. "fmt"
)

func main() {
	n := 3
	s := "A"
	aaa := 10
	var x interface{} = "00000"
	//var x interface{} = 3.141592
	i, isInt := x.(int)
	Println(i, isInt)
	f, isFloat := x.(float64)
	Println(f, isFloat)
	ss, isString := x.(string)
	Println(ss, isString)


	//goto
	Println("Start")
	goto LLLL

	switch n {
	case 1, 2:
		Println("1,2")
	case 3, 4:
		Println("3,4")
		fallthrough
	default:
		Println("defaut")
	}

	switch s {
	case "A":
		s += "Z"
		fallthrough
	case "B":
		s += "Y"
		fallthrough
	case "C":
		s += "X"
		fallthrough
	default:
		s += "W"
	}
	Println("aaaaaaaaaa == ", s)

	switch {
	case 10 < aaa && aaa < 20:
		Println("oooooo")
	case 0 < aaa && aaa <= 10:
		Println("okok")
	}


	switch x.(type) {
	case bool:
		Println("Bool")
	case string, int:
		Println("String or Int")

	}

	LLLL:
	Println("Goal")

	runDefer()

}

func runDefer() {
	defer Println("Last Run")
	Println("aaaaa")
}