package main

import (
	"fmt"
	"sync"
)

// Wait group bize goroutine'lerin bitmesini beklemek için bir yol sunar

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1) // burada kaç tane goroutine bekleyeceğimizi belirtiyoruz

	go func() {
		println("Hello from goroutine")
		wg.Done()
	}()

	wg.Wait() // burada goroutine'lerin bitmesini bekliyoruz

	fmt.Println("Hello from main")
}
