package main

import (
	"fmt"
)

func main()  {
	//3-2
	var i int
	i = 1
	fmt.Printf("hoho%d\n", i)

	a := [3]string{
		"111111",
		"222222",
		"333333",
	}
	fmt.Println(a)

	println("aaaa")

	x, y, z := 10, 20, 30

	fmt.Println(z, y, x)

	fmt.Println(one("unkkkkkkk"))

	var (
		n = 1
		b = "bbb"
		c = 3.14
	)

	fmt.Println(n, b, c)

	xxx = ngaga()
	xxx = xxx + 10
	fmt.Println(xxx)
}

var xxx = 100

func one(s string) string {
	return s + " == chinko == " + s
}

func ngaga() int {
	return xxx + 1000
}