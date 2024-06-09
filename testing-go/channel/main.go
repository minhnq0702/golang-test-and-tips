package main

import (
	"fmt"
	"sync"
)

var opponents = make(chan string)

func battle(name string, done *(chan struct{}), _wg *sync.WaitGroup) {
	fmt.Printf("===> [%s] joined to the battle of\n", name)
	select {
	case opp := <-opponents:
		fmt.Printf("The battle %s and %s\n", name, opp)
	case opponents <- name:
		fmt.Printf("Completed %v\n", &opponents)
		return
	}

	*done <- struct{}{}
	// defer _wg.Done()
}

func main() {
	lstOpp := []string{"Dog", "Cat", "Shark"}
	var wg sync.WaitGroup
	done := make(chan struct{})
	for _, opp := range lstOpp {
		fmt.Printf("Before goroutine %s\n", opp)
		// wg.Add(1)
		go battle(opp, &done, &wg)

	}
	// wg.Wait()

	select {
	case d := <-done:
		fmt.Printf("Done %v\n", d)
	case d := <-opponents:
		fmt.Printf("Opponents %v\n", d)
	}

}
