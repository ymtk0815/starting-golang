package main

import (
	. "fmt"
	"math"
)

//-- -----------------------------------------------------------メソッド
type Point struct {
	X, Y int
}

func (p *Point) Render() {
	Printf("<%d, %d>\n", p.X, p.Y)
}
func sum2() {
	p := & Point{X:5, Y:12}
	p.Render()
}

func (p *Point) Distance(dp *Point) float64 {
	dp.Y = 10
	x, y := p.X - dp.X, p.Y - dp.Y
	return math.Sqrt(float64(x*x + y*y))

}

func sum3() {
	p := & Point{X:0, Y:0}
	fuga := &Point{X:1, Y:1}
	sum3 := p.Distance(fuga)
	Println(sum3)
	Println(fuga)
}

type MyInt int

func (m MyInt) Plus(i int) int {
	return int(m) + i
}
func sum4() {
	Println(MyInt(4).Plus(2))
}