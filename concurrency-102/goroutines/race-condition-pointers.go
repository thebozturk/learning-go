package main

import (
	"fmt"
	"sync"
)

type RaceTest struct {
	Val int
}

func main() {
	raceTest := &RaceTest{} // raceTest değişkenini oluşturuyoruz ve RaceTest struct'ı ile eşleştiriyoruz

	wg := &sync.WaitGroup{} // WaitGroup oluşturuyoruz ve WaitGroup pointer'ını wg değişkenine eşleştiriyoruz
	wg.Add(1000)            // WaitGroup'un kaç goroutine bekleyeceğini belirtiyoruz
	lock := sync.Mutex{}    // mutex oluşturuyoruz

	for i := 0; i < 1000; i++ {
		go increment(raceTest, wg, &lock) // raceTest değişkenini ve wg pointer'ını increment fonksiyonuna gönderiyoruz. ayrıca mutex pointer'ını da gönderiyoruz
	}

	wg.Wait() // WaitGroup'un bitmesini bekliyoruz

	fmt.Println(raceTest)
}

func increment(rt *RaceTest, wg *sync.WaitGroup, lock *sync.Mutex) { // RaceTest pointer'ını, WaitGroup pointer'ını ve mutex pointer'ını parametre olarak alıyoruz
	// raceTest değişkeninin Val değerini 1 arttırıyoruz ve WaitGroup'un Done() fonksiyonunu çağırıyoruz
	lock.Lock() // mutex'i kilitliyoruz
	rt.Val += 1
	lock.Unlock() // mutex'i açıyoruz
	wg.Done()
}
