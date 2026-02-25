package main

func main() {
	// 管道声明为只写
	ch1 := make(chan<- int, 2)
	ch1 <- 10
	ch1 <- 20
	// <-ch1 // invalid operation: cannot receive from send-only channel chan<- int ch1 (variable of type chan<- int)

	// 声明管道为只读
	// ch2 := make(<-chan int, 2)
	// ch2 <- 10 // cannot send to receive-only channel <-chan int ch2 (variable of type <-chan int) (exit status 1)

}
