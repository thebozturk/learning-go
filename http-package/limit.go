package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strings"
	"time"
)

var counter = map[string]*Limit{
	"key": {
		count: 0,
		ttl:   time.Now(),
	},
}

// eğer /user/askljasdklj/aslkjdaksljd gibi bir istek gelirse 3 dakikada en fazla 10 istek atılabilir olsun diyebiliriz
type Limit struct {
	count int       // kaç tane istek atıldı
	ttl   time.Time // bu isteklerin sonuncusu ne zaman atıldı
}

type LimitProxy struct {
	key   string        // hangi key için limit uygulanacak
	limit int           // kaç tane istek atılabilir
	ttl   time.Duration // ne kadar sürede bir istek atılabilir
}

func NewLimitProxy(key string, limit int, ttl time.Duration) LimitProxy { // Burada LimitProxy struct'ını döndürüyoruz. LimitProxy struct'ı Proxy interface'ini implemente ediyor.
	return LimitProxy{ // LimitProxy struct'ının alanlarını dolduruyoruz.
		key:   key,
		limit: limit,
		ttl:   ttl,
	}
}

func ResetLimitHandler(c *fiber.Ctx) error {

	key := strings.TrimPrefix(c.Path(), "/limit")

	if _, ok := counter.v[key]; !ok {
		return fiber.ErrNotFound
	}

	counter.Delete(key)

	c.Response().SetStatusCode(fiber.StatusNoContent)
	return nil
}

func (p LimitProxy) Accept(key string) bool { // eğer key bu proxy için geçerliyse true döner. örnek key: user, event gibi
	return p.key == key
}

func (p LimitProxy) Proxy(c *fiber.Ctx) error {
	path := c.Path()

	if r, ok := counter.v[path]; ok && r.count >= p.limit {
		if r.ttl.After(time.Now()) {
			c.Response().SetStatusCode(fiber.StatusTooManyRequests)

			fmt.Printf("request limit reached for %s \n", path)

			return fiber.ErrTooManyRequests
		} else {
			fmt.Printf("resetting counter values \n")

			//counter[path] = &Limit{
			//	count: 0,
			//	ttl:   time.Now().Add(p.ttl),
			//}

			// thread safe version
			counter.Set(path, p.ttl)
		}
	} else if !ok {
		//counter[path] = &Limit{
		//	count: 0,
		//	ttl:   time.Now().Add(p.ttl),
		//}

		// thread safe version
		counter.Set(path, p.ttl)
	}

	if err := c.SendString("Go Turkiye - 103 Http Package"); err != nil {
		return err
	}

	//counter[path].count++

	// thread safe version
	counter.Increment(path)

	return nil
}

func (l *LimitCounter) Set(key string, ttl time.Duration) {
	l.Lock()
	l.v[key] = &Limit{count: 0, ttl: time.Now().Add(ttl)}
	l.Unlock()
}

func (l *LimitCounter) Increment(key string) {
	l.Lock()
	l.v[key].count++
	l.Unlock()
}

func (l *LimitCounter) Delete(key string) {
	l.Lock()
	delete(l.v, key)
	l.Unlock()
}
