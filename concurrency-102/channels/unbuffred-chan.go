package main

import "fmt"

// go memoryi dağıtmak yerine channel üzerinden iletişim kurar ve bu sayede memory kullanımını azaltır
// channel üzerinden iletişim kurmak için <- operatörü kullanılır

func main() {
	// channel initialization
	unbuffredChan := make(chan int) // unbuffredChannel sadace 1 değer alabilir ve bu değer okunana kadar diğer değerler bekler
	// buffredChannel ise size değeri kadar değer alabilir ve okunana kadar diğer değerler bekler

	// channel declaration
	var unbuffredChan2 chan int

	fmt.Println("unbuffredChan: ", unbuffredChan)
	fmt.Println("unbuffredChan2: ", unbuffredChan2)

	// sadece kanal üzerinden değer gönderilebilir
	go func(unbufChan <-chan int) {
		value := <-unbufChan
		fmt.Println("value: ", value)
	}(unbuffredChan)

	unbuffredChan <- 1
}
