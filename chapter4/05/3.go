package main

import (
	. "fmt"
	t "time"
)

func receive(name string, ch <-chan int)  {
	for {
		i, ok := <-ch
		if ok == false {
			break
		} else {
			Println(name, "->", i)
		}
	}
	Println(name + "is done.")

}

func main() {
	ch := make(chan int, 20)

	go receive("1st goroutine", ch)
	go receive("2nd goroutine", ch)
	go receive("3rd goroutine", ch)

	i := 0
	for i < 100 {
		ch <- i
		i++
	}
	close(ch)
	t.Sleep(3 * t.Second)

}