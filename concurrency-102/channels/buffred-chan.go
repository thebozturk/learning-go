package main

import "fmt"

func main() {
	// channel initialization
	buffredChan := make(chan int, 2) // buffredChannel ise size değeri kadar değer alabilir ve okunana kadar diğer değerler bekler
	// eğer verilen size'dan fazla değer gönderilirse deadlock oluşur

	buffredChan <- 1
	buffredChan <- 2

	// sadece kanal üzerinden değer gönderilebilir
	fmt.Println(<-buffredChan)
	fmt.Println(<-buffredChan)
}
