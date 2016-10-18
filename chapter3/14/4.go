package main

import (
	. "fmt"
	rt "runtime"
)

func main() {
	go Println("Yeah!!")

	Printf("NumCPU: %d\n", rt.NumCPU())
	Printf("NumGoroutine: %d\n", rt.NumGoroutine())
	Printf("Version: %s\n", rt.Version())




}
