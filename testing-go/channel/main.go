package main

import (
	channalbasic "channel/testbasic"
	"fmt"
	"time"
)

var opponents = make(chan string, 3)

func battle(name string, done chan struct{}) {
	time.Sleep(1 * time.Second)
	select {
	case opp := <-opponents:
		fmt.Printf("The battle %s and %s at %s\n", name, opp, time.Now())
	case opponents <- name:
		fmt.Printf("===> [%s] is waiting the battle at %s\n", name, time.Now())
	}

	fmt.Printf("<=== [%s] left the battle\n", name)

	done <- struct{}{}
}

func main() {
	channalbasic.Channel_basic()
	lstOpp := []string{"Dog", "Cat", "Shark"}

	done := make(chan struct{})
	for _, opp := range lstOpp {
		// time.Sleep(1 * time.Second)
		fmt.Printf("Before goroutine %s\n", opp)
		go battle(opp, done)
		fmt.Printf("After goroutine %s\n", opp)

	}

	count := 0
	for {
		fmt.Printf("Waiting for done %d\n", count)
		<-done
		count++
		if count == len(lstOpp) {
			break
		}
	}
	fmt.Printf("Done\n")

}
