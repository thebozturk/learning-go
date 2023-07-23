package main

import "fmt"

// channels: channelslar goroutine'lar arasında iletişim kurmak için kullanılır. Bir kanal oluşturmak için make fonksiyonu kullanılır.
// Kanalın tipi kanalda taşınacak verinin tipidir.
// goroutinelerde, goroutine alan fonksiyon hiçbir şey return etmez. Bunun yerine kanal üzerinden değer döndürür. cünkü goroutine'lerin çalışma zamanı bilinmez. Bu yüzden kanal üzerinden değer döndürmek daha mantıklıdır.
// eğer 2. goroutine olmazsa deadlock oluşur. deadlock oluşmaması için kanalın içine değer gönderilmesi ve alınması aynı anda olmamalıdır.

type Person struct {
	name string
	age  int
}

func displayPerson(p Person, myChannel chan string) {
	myChannel <- fmt.Sprintf("Name: %s, Age: %d", p.name, p.age)
}
func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go displayPerson(Person{"John", 25}, chan1)
	go displayPerson(Person{"Jane", 30}, chan2)

	fmt.Println(<-chan1)
	fmt.Println(<-chan2)
}
