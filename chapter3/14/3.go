package main

import (
	. "fmt"
)

func sub(s string) {
	for {
		Println("subb", s)
	}
}

func main() {
	go sub("11111111")
	go sub("2222222222")

	for {
		Println("mainnnnnnnnnnnnnnnnn")
	}

}