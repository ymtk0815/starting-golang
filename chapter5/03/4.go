package main

import (
	. "fmt"
)

type User struct {
	Id   int
	Name string
}

func NewUser(id int, name string) *User {
	u := new(User)
	u.Id = id
	u.Name = name
	return u
}

func sum41() {
	user := NewUser(1, "Taro")
	Println(*user)
}

type Point struct {
	X, Y int
}

func (p *Point) ToString() string {
	return Sprintf("[%d, %d]", p.X, p.Y)
}
func sum42() {
	f := (*Point).ToString
	Println(f(&Point{X: 7, Y: 11}))
	Println(((*Point).ToString)(&Point{X: 11, Y: 33}))

	pp := &Point{X: 2, Y: 3}
	ff := pp.ToString
	Println(ff())
}

func main() {
	sum41()
	Println("-- -----------------------------------------\n")
	sum42()
	Println("-- -----------------------------------------\n")

}
