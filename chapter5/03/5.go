package main

import (
	. "fmt"
)

type Point struct {
	X, Y int
}

func sum51() {
	ps := make([]Point, 5)
	ps[0].X = 10
	ps[0].Y = 20
	for _, p := range ps {
		Println(p.X, p.Y)
	}
}

type Points []*Point

func (ps Points) ToString() string {
	str := ""
	for _, p := range ps {
		if str != "" {
			str += ","
		}
		if p == nil {
			str += "<nil>"
		} else {
			str += Sprintf("[%d, %d]", p.X, p.Y)
		}
	}
	return str
}
func sum52() {
	ps := Points{}
	ps = append(ps, &Point{X: 1, Y: 2})
	ps = append(ps, nil)
	ps = append(ps, &Point{X: 3, Y: 4})
	Println(ps.ToString())
}

type User struct {
	Id   int
	Name string
}

func sum53() {
	m1 := map[User]string{
		{Id: 1, Name: "Taro"}: "Tokyo",
		{Id: 2, Name: "Jiro"}: "Osaka",
	}
	Println(m1)
	m2 := map[int]User{
		1: {Id: 3, Name: "Saburo"},
		2: {Id: 4, Name: "Shiro"},
	}
	Println(m2)
}

func sum54() {
	ms := map[int][]string{
		1: {"A", "B", "C"},
	}
	nm := map[int]map[int]string{
		1: {
			1: "Apple",
			2: "Banana",
			3: "Cherry",
		},
	}
	Println(ms)
	Println(nm)
}

func main() {
	sum51()
	Println("----------------------\n")
	sum52()
	Println("----------------------\n")
	sum53()
	Println("----------------------\n")
	sum54()
	Println("----------------------\n")
	Println("----------------------\n")

}
