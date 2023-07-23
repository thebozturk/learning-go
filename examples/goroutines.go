package main

import (
	"fmt"
	"sync"
)

// goroutineler sayesinde birden fazla işlemi aynı anda yapabiliriz. (concurrency)
// eş zamanlı yapılan işlemlerin her birine goroutine denir.
// goroutine'lerin çalışması için main fonksiyonunun bitmesi gerekir.

var wg sync.WaitGroup // goroutine'lerin bitmesini beklemek için kullanıyoruz.

func main() {
	wg.Add(1) // 1 tane goroutine beklediğimizi belirtiyoruz.
	go printX()
	wg.Wait() // goroutine'lerin bitmesini bekliyoruz.
	printY()
}

func printX() {
	for i := 0; i < 10; i++ {
		fmt.Println("x:", i)
	}
	wg.Done() // burada işlemin tamamlandıgını belirtiyoruz. goroutine'lerden biri bittiğinde wg.Add(1) ile eklediğimiz sayıyı 1 azaltıyoruz.
}

func printY() {
	for i := 0; i < 10; i++ {
		fmt.Println("y:", i)
	}
}
