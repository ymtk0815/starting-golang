package main

import (
	. "fmt"
)

func main() {
	ch := make(chan string, 3)
	ch <- "Apple"
	Println(len(ch))
	ch <- "Bananao"
	ch <- "Cherry"
	Println(len(ch))
	Println("------------------\n\n\n\n")

	aa := make(chan string)
	Println(cap(aa))

	bb := make(chan string, 3)
	Println(cap(bb))
	Println(len(bb))

	cc := make(chan int, 3)
	cc <- 1
	cc <- 2
	cc <- 3
	close(cc)
	Println(<-cc)
	Println(<-cc)
	Println(<-cc)
	Println(<-cc)
	Println(<-cc)
	Println(<-cc)
	Println("------------------\n\n\n\n")

	dd := make(chan int)
	close(dd)
	ee, ok := <-dd
	Println(ee)
	Println(ok)


	Println("------------------\n\n\n\n")

	ff := make(chan int, 3)
	ff <- 1
	ff <- 2
	ff <- 3
	close(ff)
	var (
		i int
		okk bool
	)
	i, okk = <-ff
	Println(i)
	Println(okk)
	i, okk = <-ff
	Println(i)
	Println(okk)
	i, okk = <-ff
	Println(i)
	Println(okk)
	i, okk = <-ff
	Println(i)
	Println(okk)


}
