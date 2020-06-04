package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"runtime/debug"
	"time"
)

type Foo struct {
	a int64
}

func main() {
	debug.SetGCPercent(-1)
	var ms runtime.MemStats

	runtime.ReadMemStats(&ms)
	fmt.Printf("allocation: %f Mb - %d\n", float32(ms.HeapAlloc)/float32(1024*1024), ms.HeapObjects)

	for i := 0; i <= 1000000; i++ {
		f := NewFoo(i)
		_ = fmt.Sprintf("%d", f.a)
	}

	runtime.ReadMemStats(&ms)
	fmt.Printf("allocation: %f Mb - %d\n", float32(ms.HeapAlloc)/float32(1024*1024), ms.HeapObjects)

	runtime.GC()
	time.Sleep(time.Second)

	runtime.ReadMemStats(&ms)
	fmt.Printf("Final allocation: %f Mb - %d\n", float32(ms.HeapAlloc)/float32(1024*1024), ms.HeapObjects)

	runtime.GC()
	time.Sleep(time.Second)

	runtime.ReadMemStats(&ms)
	fmt.Printf("2 Force allocation: %f Mb - %d\n", float32(ms.HeapAlloc)/float32(1024*1024), ms.HeapObjects)
}

//go:noinline
func NewFoo(i int) *Foo {
	newFoo := &Foo{
		a: rand.Int63n(100),
	}
	runtime.SetFinalizer(newFoo, func(f *Foo) {
		_ = fmt.Sprintf("%+v %d has been garbage collected\n", f, i)
	})
	return newFoo
}
