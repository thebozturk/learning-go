package main

import "fmt"

func main() {
	// channel initialization
	buffredChan := make(chan int, 2)

	go func(buffChan chan int) { // buffChan burada bir pointer ve bu sayede değerlerini değiştirebiliriz
		for {
			value := <-buffChan
			fmt.Println("value: ", value)
		}
	}(buffredChan)

	buffredChan <- 1
	buffredChan <- 2

	fmt.Println(<-buffredChan)
}
