package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = "Init"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	fmt.Println(reflect.TypeOf(f))
}
