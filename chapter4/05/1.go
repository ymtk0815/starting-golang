package main

import (
	. "fmt"
)

func receiver(ch <- chan int) {
	for {
		i := <- ch
		Println(i)
	}
}


func main() {
	//aa := make(chan int, 10)
	//aa <- 5
	//i := <-aa
	//Printf("%v", i)

	ch := make(chan int)
	go receiver(ch)
	i := 0
	for i < 10000 {
		ch <- i
		Println("aaaaaa")
		i++
	}
}
