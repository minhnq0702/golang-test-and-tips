package main

import "fmt"

func main() {
	a := int64(1)

	switch a {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		// do no thing
		fallthrough
	default:
		fmt.Println("This is default!!!!")

	}
}
