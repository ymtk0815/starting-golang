package main

import (
	. "fmt"
)

// fafaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
type Point struct {
	X int
	Y int
}

func hoge1() {
	pt := Point{}
	pt.X = 10
	pt.Y = 10
	Println(pt.X)
	Println(pt.Y)
	pt2 := Point{10, 22}
	Println(pt2)
}

//bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb
type Person struct {
	ID    int
	name  string
	busyo string
}

func hoge2() {
	p := Person{busyo: "eigyo", name: "yamashita", ID: 10}
	Println(p)

}

//構造体を含む構造体aaaaaaaaaaaaaaaaaaaaaaa
type Feed struct {
	Name   string
	Amount int
}
type Animal struct {
	Name string
	Feed Feed
}

func hoge3() {
	a := Animal{
		Name: "Monkey",
		Feed: Feed{
			Name:   "Banana",
			Amount: 10,
		},
	}
	Println(a.Name)
	Println(a.Feed.Name)
	Println(a.Feed.Amount)
	a.Feed.Amount = 200
	Println(a.Feed.Amount)
}

//faaaaaaaaaaafaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
type T0 struct {
	Name1 string
}
type T1 struct {
	T0
	Name2 string
}
type T2 struct {
	T1
	Name3 string
}

func hoge4() {
	t := T2{
		T1: T1{
			T0: T0{
				Name1: "X",
			},
			Name2: "Y",
		},
		Name3: "Z",
	}
	Println(t.Name3)
	Println(t.Name2)
	Println(t.Name1)
}

//aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
type Base struct {
	Id    int
	Owner string
}
type A struct {
	Base /* 共通 */
	Name string
	Area string
}
type B struct {
	Base
	Title  string
	Bodies []string
}

func hoge5() {
	a := A{
		Base: Base{17, "Taro"},
		Name: "Jiro",
		Area: "Tokyo",
	}
	b := B{
		Base:   Base{99, "Hanako"},
		Title:  "no title",
		Bodies: []string{"hoge", "fuga"},
	}
	Println(a.Id)
	Println(a.Area)
	Println(a.Name)
	Println(a.Owner)
	Println(b.Id)
	Println(b.Owner)
	Println(b.Title)
	Println(b.Bodies)
}


func main() {
	hoge1()
	Println("---------------------\n")
	hoge2()
	Println("---------------------\n")
	hoge3()
	Println("---------------------\n")
	hoge4()
	Println("---------------------\n")
	hoge5()
	Println("---------------------\n")

}
