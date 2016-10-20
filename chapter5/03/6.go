package main

import (
	. "fmt"
	"reflect"
)

type User struct {
	Id   int    "ID"
	Name string "名前"
	Age  int    "年齢"
}

func sum61() {
	u := User{
		Id:   1,
		Name: "Taro",
		Age:  10,
	}
	t := reflect.TypeOf(u)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		Println(f.Name, f.Tag)
	}

}

func main() {
	sum61()

}
