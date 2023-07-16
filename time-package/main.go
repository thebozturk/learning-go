package main

import (
	"fmt"
	"time"
)

/*
Unix time (Unix zamanı), 1 Ocak 1970 00:00:00 UTC'den bu yana geçen saniye sayısını ifade eder.
Unix zamanı, zamanı ifade etmek için kullanılan en yaygın yöntemdir.

9 Eylük 2001 (09/09/2001) saat 04:46:40 itibari ile Unix zaman 10 basamaklı bir sayıya ulaşmıştır.
Ve 20 Kasım 2286 (20/11/2286) saat 19:46:40 itibari ile Unix zaman 11 basamaklı bir sayıya ulaşacaktır.

*/

func main1() {
	fmt.Print("Şu anki unix zamanı: ", time.Now().Unix())
	// Zamanı bi saniye bekletelim. Çünkü zamanı alırken bir saniye geçmiş olabilir. :)
	time.Sleep(time.Second)
	fmt.Print("\nŞu anki unix zamanı: ", time.Now().Unix())
}

func main() {
	// Gün ay yıl bilgisi almak için time.Now().Day(), time.Now().Month(), time.Now().Year() fonksiyonlarını kullanabiliriz.
	fmt.Println("Bugün günlerden", time.Now().Day(), time.Now().Month(), time.Now().Year())

	// Yer bilgisi almak için time.Now().Location() fonksiyonunu kullanabiliriz.
	fmt.Println("Yer bilgisi:", time.Now().Location())
}
