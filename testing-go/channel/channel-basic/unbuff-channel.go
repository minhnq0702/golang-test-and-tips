package channalbasic

func TestUnBuffChannel() {
	c := make(chan int)
	c <- 1 // * cause deadlock and PANIC
}
