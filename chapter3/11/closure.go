package main

import (
	"fmt"
)

func later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func integers() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func fafafa() func(int) int {
	i := 20
	return func(ii int) int {
		i += ii
		return i
	}

}

func main() {
	//fmt.Printf("ddddddddddddd")

	f := later()
	fmt.Println(f("Golang"))
	fmt.Println(f("is"))
	fmt.Println(f("awesome!"))

	ints := integers()

	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	fa := fafafa()

	fmt.Println(fa(10000))
	fmt.Println(fa(1000))
	fmt.Println(fa(100))

}