package main

import (
	. "fmt"
)

func main() {
	type MyInt int

	var n1 MyInt = 5
	n2 := MyInt(7)
	Println(n1)
	Println(n2)
	Println("\n---------------\n")

	type (
		IntPair     [2]int
		Strings     []string
		AreaMap     map[string][2]float64
		IntsChannel chan []int
	)

	pair := IntPair{
		1,
		2,
	}
	strs := Strings{
		"Apple",
		"Banana",
		"Cherry",
	}
	amap := AreaMap{
		"Tokyo": {11.1, 22.2},
		"Osaka": {33.3, 44.4},
	}
	ich := make(IntsChannel)
	Println(pair)
	Println(strs)
	Println(amap)
	Println(ich)
	Println("\n---------------\n")

	n := Sum(
		[]int{1, 2, 3, 4, 5},
		func(i int) int {
			return i * 2
		},
	)
	Println(n)

	Println("\n---------------\n")
	type T0 int
	type T1 int
	t0 := T0(5)
	i0 := int(t0)
	Printf("%T\n", t0)
	Printf("%T\n", i0)
	Println(t0)
	Println(i0)

	t1 := T1(8)
	i1 := int(t1)
	t1 = T1(i1)
	Println(t1)
}

type Callback func(i int) int

func Sum(ints []int, callback Callback) int {
	var sum int
	for _, i := range ints {
		sum += i
	}
	return callback(sum)
}
