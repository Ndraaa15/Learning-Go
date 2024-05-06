package main

func main() {

	// Buffered Channel
	// Channel that can store limited value
	ch1 := make(chan int, 2)

	// Unbuffered Channel
	ch2 := make(chan int)

	/*
		ch <- 1
		n := <-ch
		fmt.Println(n)

		Code above will cause deadlock because channel is not buffered
		And also when we send or recieve data it must be at the same time so wee need to use goroutine
		Channel is like portal so after we go into a portal and simustaneously we go out from the portal
	*/

	/*
		go func() {
			ch <- 1
		}()
		n := <-ch
		println(n)

		Code above will work properly because already using goroutine
		This make the channel can send and recieve data at the same time (simultaneously)
		This is called asynchronous
	*/

	ch1 <- 1
	num1 := <-ch1
	println(num1)

	go func() {
		ch2 <- 2
	}()

	num2 := <-ch2

	print(num2)
}
