package main

import "fmt"

var opponents = make(chan string)

func battle(name string, done chan struct{}) {
	select {
	case opp := <-opponents:
		fmt.Printf("The battle %s and %s", name, opp)
	case opponents <- name:
		fmt.Printf("%v", opponents)
	}

	done <- struct{}{}

}

func main() {
	lstOpp := []string{"Dog", "Cat", "Shark"}
	done := make(chan struct{})
	for _, opp := range lstOpp {
		fmt.Printf("Before goroutine %s\n", opp)
		go battle(opp, done)
	}
}
