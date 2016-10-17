package main

import (
	"fmt"
)

const (
	XXX = 1
	YYY = 2
	ZZZ = 3

	CCC = 10
	DDD
	EEE = "aaaaaaaaa"
	FFF
)

const AAA = 10

const (
	GGG = 10 + iota
	HHH = 20 + iota
)

const (
	III = iota
	JJJ
)

func one(a int) int {
	return a + AAA
}

func main() {
	fmt.Println(XXX, YYY, ZZZ)
	fmt.Println(one(10))
	fmt.Println(CCC, DDD, EEE, FFF)
	fmt.Println(GGG, HHH)
	fmt.Println(III, JJJ)
}