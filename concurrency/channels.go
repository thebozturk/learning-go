package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// Bir değer göndermek için <- operatörü kullanılır.
	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	// Kanal oluşturmak için make fonksiyonu kullanılır.
	// Kanalın türü kanalı geçen değerlerin türüdür.
	c := make(chan int)
	// Kanalı kullanmak için <- operatörü kullanılır.
	// Bu örnekte sum fonksiyonu goroutine olarak çalıştırılır.
	go sum(s[:len(s)], c)
	// Kanaldan değer almak için <- operatörü kullanılır.
	// Bu örnekte kanaldan gelen değer x değişkenine atanır.
	x := <-c
	// Kanaldan gelen değer ekrana yazdırılır.
	fmt.Println(x)
}
