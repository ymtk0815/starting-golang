package main

import (
	"fmt"
)

func main() {

	hello()

	aa, bb := div(10, 3)
	fmt.Printf("%v == %v\n", aa, bb)

	//戻り値の破棄
	cc, _ := div(20, 3)
	_, dd := div(20, 3)
	fmt.Printf("%v |||| %v\n", cc, dd)

	fmt.Println(doSomething())


	eee := func(x, y int) int {
		return x + y
	}
	fmt.Println(eee(1, 3))
	fmt.Printf("%T\n", func(x, y int) int {return x + y})

	callFunc(func() {
		fmt.Println("aaaaaaaaaaa")
	})

}

//戻り値のない関数
func hello() {
	fmt.Println("Hello, hogheogae")
	return
}

//複数の戻り値
func div(a, b int) (int, int) {
	q := a * b
	p := a % b
	return q, p
}

//戻り値を表す変数
func doSomething() (x, y int) {
	y = 5
	return
}

func callFunc(f func()) {
	fmt.Println("vvvvvvvvvvvvvvv")
	f()
}



