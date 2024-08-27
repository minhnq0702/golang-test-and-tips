package channalbasic

func TestBuffChannel() {
	c := make(chan int, 1)
	c <- 1 // * do not cause deadlock and continue running to below
}
