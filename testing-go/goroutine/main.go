package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func atomicCount(id int, counter *int32, wg *sync.WaitGroup) {
	_ = id
	defer wg.Done() // * decrease wg by 1
	// defer mut.Unlock()
	// mut.Lock()
	// *counter += 1
	atomic.AddInt32(counter, 1)
}

func main() {
	number_of_counter := 100000000
	var counter int32 = 0
	fmt.Printf("starting\n")
	var wg sync.WaitGroup

	wg.Add(number_of_counter) // * add number of task need to do into WG

	start_at := time.Now()
	for id := range number_of_counter {
		go atomicCount(id, &counter, &wg)
	}

	wg.Wait() // * wait until WG count to zero
	eslapsed := time.Since(start_at)
	fmt.Printf("Done %d with %vms\n", counter, eslapsed.Milliseconds())
}
