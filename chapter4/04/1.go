package main

import (
	. "fmt"
)

func main() {

	aa := make(map[int]string)

	aa[1] = "US"
	aa[81] = "Japan"
	aa[86] = "China"
	Println(aa)

	bb := make(map[string]string)
	bb["Yamada"] = "Taro"
	bb["Sato"] = "Hanako"
	//bb["Yamada"] = "Jiro"
	Println(bb)

	cc := make(map[float64]int)
	cc[1.02] = 1
	cc[1.03] = 2
	Println(cc)

	dd := map[int]string{
		1: "Taro",
		2: "Hanako",
		3: "Jiro",
	}
	Println(dd)

	ee := map[int][]int{
		1: []int{1},
		2: []int{1, 2},
		3: []int{1, 2, 3},
	}
	Println(ee)

	ff := map[int][]int{
		1: {1},
		2: {1, 2},
		3: {1, 2, 3},
	}
	Println(ff)

	gg := map[int]map[float64]int{
		1: {3.14: 11},
	}
	Println(gg)

	hh := map[int]string{
		1: "A",
		2: "B",
		3: "C",
		4: "D",
	}
	ii := hh[1]
	jj := hh[9]
	Println(hh)
	Println(ii)
	Println(jj)

	kk := hh[7]
	mm, error := hh[9]
	ll, error := hh[1]
	Println(error)
	Println(kk)
	Println(ll)
	Println(mm)
	Println(error)

	if _, ok := hh[1]; ok {
		Println(ok)
		Println("okkkkkkkkkkkk")
	}

	nn := map[int][]int {
		1: {1},
		2: {1, 2},
		3: {1, 2, 3},
	}
	s := nn[10]
	if s != nil {
		Println("bbbbbbbbbbb")
	} else {
		Println("なんもnai")
	}



}
