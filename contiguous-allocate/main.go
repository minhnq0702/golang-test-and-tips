package main

import (
	"fmt"
	"unsafe"
)

func main() {
	const a = 1111111111111111111

	// x := [7]int{a, 1, 33333, 4, 6, 7, 7}
	// * or
	x := []int{1, 2, 3, 4}

	unsafePointer := unsafe.Pointer(&x[0])
	fmt.Printf("Starting addr %p\n", unsafePointer)
	step := unsafe.Sizeof(int(0))
	x = append(x, a)
	for i := range x {
		addr_n := unsafe.Add(unsafePointer, int(step)*i)
		fmt.Printf("addr: %p, val: %d with size of %d\n ", addr_n, *(*int)(addr_n), step)
	}
}
