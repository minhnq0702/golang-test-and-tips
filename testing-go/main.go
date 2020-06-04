package main

import (
	"bytes"
	"crypto/subtle"
	"fmt"
	"html/template"
	"os"
	"time"

	"github.com/patrickmn/go-cache"
)

var c *cache.Cache

func addCache(step int) {
	// time.Sleep(1 * time.Second)
	err := c.Add("x", int64(1), cache.DefaultExpiration)
	if err != nil {
		fmt.Println("Key exists, so increase it", step)
		_, err = c.IncrementInt64("x", 1)
		if err != nil {
			fmt.Println("error===>", err.Error())
		}
	}
}

func getCache(c *cache.Cache, step int) {
	value, ok := c.Get("x")
	if ok {
		fmt.Printf("value of step %d: %d\n", step, value)
	} else {
		fmt.Println("not exits", ok)
	}
}

func testNil(a *int, b int) {
	fmt.Println(a, b)
}

func myFunc(a *int64) {
	*a++
}

func TestDefer(df bool) {
	defer fmt.Println("This is defer")
	if df {
		fmt.Println("Get before defer")
		return
	}
}

func main() {
	// c = cache.New(2*time.Minute, 5*time.Minute)
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		go addCache(i)
	// 		go getCache(c, i)
	// 	}
	// }()
	// time.Sleep(5 * time.Second)

	// value, _ := c.Get("x")
	// fmt.Printf("value of final step %d\n", value)

	a := map[string]map[string]string{
		"a": map[string]string{
			"b": "c",
		},
	}
	fmt.Println(a["a"]["b"])

	testNil(nil, 1)

	text := "Hello {{.name}}\nI'm {{.myName}}\nAnd at {{index .slice 1}}\n"
	params := map[string]interface{}{
		"name":   "MINH",
		"myName": "TADA",
		"slice":  []string{"1", "2", "3"},
	}
	temp, _ := template.New("Test").Parse(text)
	temp.Execute(os.Stdout, params)

	temp, _ = template.New("TestNew").Parse("This is slice {{index . 2}}")
	buf := new(bytes.Buffer)
	temp.Execute(buf, []int{4, 5, 6})
	fmt.Println(buf.String())

	x := int64(10)
	myFunc(&x)
	fmt.Println(x)

	myMap := make(map[string]string)
	mySlice := make([]int, 10)

	myMap["a"] = "b"
	fmt.Println(myMap, mySlice)

	mySlice = append(mySlice, 1)
	fmt.Println(myMap, mySlice)

	var myA [10]int
	fmt.Println(myA)

	TestDefer(true)

	var myDict = make(map[int64][]int64)
	array := []int64{1, 2, 3}
	for _, i := range array {
		myDict[i] = append(myDict[i], 10)
	}

	for key, val := range myDict {
		fmt.Println(key, "===", val)
	}

	// {1582e56614788371f82f8948882bd0cf,cbf425f230afaa5961b98a94265ec77e}
	// 306d0aec87f2d8ad2eeb30436ae2b5eb

	// 6a0a71e6ab9a9f822c33c14ffbd1f4d7,df277614857b3eea7bd69b0ca917be11
	// 6a0a71e6ab9a9f822c33c14ffbd1f4d7,df277614857b3eea7bd69b0ca917be11
	// 6a0a71e6ab9a9f822c33c14ffbd1f4d7,df277614857b3eea7bd69b0ca917be11,4a31830b12c6e73f2ff413d0905dd69d
	// df277614857b3eea7bd69b0ca917be11
	com := subtle.ConstantTimeCompare([]byte("df277614857b3eea7bd69b0ca917be11"), []byte("df277614857b3eea7bd69b0ca917be11"))
	fmt.Println("===>", com)

	newTime := "2020-06-01 00:00:00+07:00"
	_, err := time.Parse("2006-01-02 15:04:05+07:00", newTime)
	println(err.Error())

}
