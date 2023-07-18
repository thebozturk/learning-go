package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
		// buraya i veriyoruz çünkü i değişkeni her seferinde değişiyor eğer vermezsek her seferinde 10 yazacak cünkü i değişkeni döngü bittiğinde 10 oluyor bu sayede goroutine'ler çalıştığında i değişkeni 10 oluyor
		// clojure'lerde bu şekilde çalışıyor yani içerideki fonksiyon dışarıdaki değişkenleri kullanabiliyor
	}
	time.Sleep(time.Second * 3)
}
