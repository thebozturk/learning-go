package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// basit bir goroutine örneği
	go xFunc()
	// goroutine'lerin çalışması için main thread'in beklemesi gerekiyor
	// aksi takdirde main thread sonlanır ve goroutine'ler de sonlanır. goroutine kullanıdıgımızda ise yeni bir thread oluşturulur ve bu thread main thread sonlandıktan sonra da çalışmaya devam eder
	time.Sleep(100 * time.Millisecond)

	// fonksiyonun basında go yoksa bu main thread'de çalışır eğer varsa yeni bir thread'de çalışır
	xFunc()

	// goroutine bir thread gibidir fakat bir thread'den daha hafiftir.
	// işlemci seviyesinde bir thread'den daha hafif olması sebebiyle bir thread'den daha fazla goroutine çalıştırabiliriz

	// basit bir paralelism örneği
	runtime.GOMAXPROCS(1)
	go xFunc()
	time.Sleep(100 * time.Millisecond)
}

func xFunc() {
	// go routine
	for l := byte('a'); l <= byte('z'); l++ {
		fmt.Println(string(l))
	}

	// paralelism
	for l := byte('A'); l <= byte('Z'); l++ {
		go fmt.Println(string(l))
	}
}
