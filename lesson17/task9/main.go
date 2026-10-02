package main

func main() {
	ch := make(chan int, 6)

	for i := 1; i < 3; i++ {
		ch <- i
	}

	close(ch)
	// panic: close of closed channel - пытаемся закрыть уже закрытый канал,
	// канал закрывается всего 1 раз
	// close(ch)




	// panic: send on closed channel - пытаемся отправить в уже закрытый канал,
	// ch <- 6
}
