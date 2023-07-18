package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// race condition: iki veya daha fazla goroutine'in aynı anda bir değişkene erişmesi ve bu değişkeni değiştirmesi durumudur
// bu durumda bir goroutine'in değiştirdiği değer diğer goroutine tarafından değiştirilmeden kullanıldığında beklenmeyen sonuçlar ortaya çıkabilir

func raceExample() {
	wg := sync.WaitGroup{}
	wg.Add(2) // burada kaç tane goroutine bekleyeceğimizi belirtiyoruz

	// shared value
	val := 0

	go func() {
		for i := 0; i < 100000000; i++ {
			val++
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100000000; i++ {
			val++
		}
		wg.Done()
	}()

	wg.Wait() // burada goroutine'lerin bitmesini bekliyoruz

	fmt.Println(val)
}

func raceExampleFixed() {
	// bu problemleri mutex ile çözebiliriz
	// mutex bize bir değişkene aynı anda sadece bir goroutine'in erişmemizi sağlar. diğer goroutine'ler bu değişkene erişmek istediğinde bekler.
	// yani eş zamanlı olarak sadece bir goroutine bu değişkene erişebilir ve değiştirebilir
	wg := sync.WaitGroup{}
	wg.Add(2)

	lock := sync.Mutex{} // mutex oluşturuyoruz

	//shared value
	val := 0

	go func() {
		for i := 0; i < 100000000; i++ {
			lock.Lock() // mutex'i kilitliyoruz ve diğer goroutine'lerin bu değişkene erişmesini engelliyoruz
			val++
			lock.Unlock() // mutex'i açıyoruz ve diğer goroutine'lerin bu değişkene erişmesine izin veriyoruz
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100000000; i++ {
			lock.Lock()
			val++
			lock.Unlock()
		}
		wg.Done()
	}()

	wg.Wait()

	fmt.Println(val)
}

func raceExampleFixedWithAtomic() {
	// mutex kullanmak yerine atomic paketini de kullanabiliriz
	// atomic paketi ile bir değişkeni eş zamanlı olarak değiştirebiliriz
	// mutex'e göre daha performanslıdır
	// atomic'de primitive tipler ile çalışır. yani int, float, string gibi tipler ile değil int32, int64, uint32, uint64, uintptr gibi tipler ile çalışır
	wg := sync.WaitGroup{}
	wg.Add(2)

	//shared value
	var val int32 = 0

	go func() {
		for i := 0; i < 100000000; i++ {
			atomic.AddInt32(&val, 1) // val değişkenine 1 ekliyoruz
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100000000; i++ {
			atomic.AddInt32(&val, 1)
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(val)
}

func main() {
	//raceExample()
	//raceExampleFixed()
	raceExampleFixedWithAtomic()
}
