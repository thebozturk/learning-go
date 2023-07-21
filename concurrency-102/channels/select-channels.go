package main

import "fmt"

func main() {
	chan1 := make(chan int, 1)

	chan1 <- 1

	select { // select ile birden fazla kanalı dinleyebiliriz
	case c1Val := <-chan1: // eğer kanal üzerinden değer gelirse bu case çalışır
		fmt.Println("c1Val: ", c1Val)
	}
}
