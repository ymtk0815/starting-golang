package main

import (
	. "fmt"
	"./foo"
)

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{
		Message: "エラーが発生しました",
		ErrCode: 1234,
	}
}

func sum11() {
	err := RaiseError()
	Println(err.Error())

}

type Stringify interface {
	ToString() string
}
type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return Sprintf("%s(%d)", p.Name, p.Age)
}
func (p Person) ToString2() string {
	return Sprintf("bbbbb%s(%d)", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return Sprintf("[%s %s]", c.Number, c.Model)
}

func pppp(s Stringify) {
	Println(s.ToString(), "aaaaaaaaaaaaaaaaaaaa")
}
func sum12() {
	vs := []Stringify{
		&Person{Name: "Taro", Age: 21},
		&Car{Number: "12-34", Model: "PS512"},
	}
	for _, v := range vs {
		Println(v.ToString())
	}

	pppp(&Person{Name:"Hanako", Age:10})
	Println((&Person{Name:"Hanako", Age:10}).ToString())
	Println((Person{Name:"Hanako", Age:10}).ToString2())
}

func sum13() {
	t := foo.NewI()
	Println(t.Method1())
	//Println(t.method2())
}



func main() {
	sum11()
	sum12()
	sum13()


}
