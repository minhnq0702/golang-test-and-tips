package main

import (
	"fmt"
	"reflect"
)

type A struct {
	X int
	Y int
}

type B struct {
	X int
	Y int
}

func main() {
	a := A{
		X: 1,
		Y: 1,
	}

	b := B{}

	b = B(a)

	fmt.Println(b)

	structVal := reflect.ValueOf(a)
	fmt.Println("==Reflect", structVal.FieldByName("X").Int())
	fmt.Println("==Reflect", reflect.Indirect(structVal))

	test := map[string]interface{}{"a": "a"}
	test2, ok := test["b"].(string)
	fmt.Println("====>", len(test2), ok)
}
