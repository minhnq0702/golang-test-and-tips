package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type BaseFace interface {
	HelloWorld(string)
}

type Model1 struct{}
type Model2 struct{}

type StructToZero struct {
	Name string
}

func (model Model1) HelloWorld(msg string) {
	fmt.Println("Model 1 say hello", msg)
}

func (model Model2) HelloWorld(msg string) {
	fmt.Println("Model 2 say hello", msg)
}

func TestModel() {
	model1 := new(Model1)
	model2 := new(Model2)

	lstModel := []interface{}{model1, model2}

	for i, model := range lstModel {
		// model.HelloWorld(strconv.Itoa(i))
		ProcessIt(model.(BaseFace), strconv.Itoa(i))
	}

	// ProcessIt(model1, "VN")
	// ProcessIt(model2, "USA")

	toCompare := StructToZero{Name: ""}
	CompareZeroStruct(toCompare)

	// Method expression
	var testFunc func(BaseFace, string) = BaseFace.HelloWorld
	testFunc(*model2, "Test method expression")
	// fmt.Println(testFunc)
}

func main() {
	var test *string
	if test == nil {
		fmt.Println("nil ne")
	}

	var test2 int64
	fmt.Printf("==> %v \n", test2)

	today := time.Now().UTC()
	previous := today.AddDate(0, 0, -91)
	fmt.Println(previous)

}

func ProcessIt(model BaseFace, msg string) {
	model.HelloWorld(msg)
}

func CompareZeroStruct(toCompare StructToZero) {
	// ZeroStruct := struct{}{}
	if reflect.DeepEqual(toCompare, StructToZero{}) {
		fmt.Println("This is zero value")
	} else {
		fmt.Println("Not Zero")
	}
}
