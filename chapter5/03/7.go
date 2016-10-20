package main

import (
	. "fmt"
	"encoding/json"
)

type User struct {
	Id int `json:"user_id"`
	Name string `json:"user_name"`
	Age int `json:"user_age"`
}

func sum71() {
	u := User{Id:1, Name:"Taro", Age:32}
	bs, _ := json.Marshal(u)
	Println(string(bs))
}

func main() {
	sum71()
}