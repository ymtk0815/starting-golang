package main

import (
	. "fmt"
)

//-- -----------------------------------------------------------new
type Person struct {
	Id   int
	Name string
	Area string
}

func sum1() {
	p := new(Person)
	p.Id = 10
	p.Name = "Jiro"
	p.Area = "Tokyo"
	Println(p.Id)
	Println(p.Name)
	Println(p.Area)
}

type IntPair [2]int

func (ip IntPair) First() int {
	return ip[0]
}
func (ip IntPair) Last() int {
	return ip[1]
}

type Strings []string

func (s Strings) Join(d string) string {
	sum := ""
	for _, v := range s {
		if sum != "" {
			sum += d
		}
		sum += v
	}
	return sum
}
func plus1() {
	ip := IntPair{1, 2}
	Println(ip.First())
	Println(ip.Last())
	Println(Strings{"A", "B", "C"}.Join(","))
}

func main() {
	sum1()
	Println("-- -----------------------------------------\n")
	sum2()
	Println("-- -----------------------------------------\n")
	sum3()
	Println("-- -----------------------------------------\n")
	sum4()
	Println("-- -----------------------------------------\n")
	plus1()
	Println("-- -----------------------------------------\n")

}
