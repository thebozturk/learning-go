package main

import "fmt"

func main() {
	// go bunu alıp x bir zamanda çalıştıracak
	go func() {
		println("Hello from goroutine")
	}()
	fmt.Println("Hello from main")
}
