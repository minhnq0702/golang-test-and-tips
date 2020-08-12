package main

import "fmt"

func main() {
	a := 2
	switch a {
	case 2:
	case 3:
	default:
		fmt.Println("Invalid")
	}

	swithFallthourgh()
}

func swithFallthourgh() {
	a := int64(1)

	switch a {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		// Fall all to default
		fallthrough
	default:
		fmt.Println("This is default!!!!")

	}
}
