package main

import (
	"fmt"
)

func main() {

	var m map[string]string = make(map[string]string)
	m["who"] = "world"

	fmt.Println("Hello", m["who"], "again!")
}
