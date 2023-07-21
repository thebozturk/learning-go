package main

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

type Proxy interface {
	Accept(key string) bool   // eğer key bu proxy için geçerliyse true döner. örnek key: user, event gibi
	Proxy(c *fiber.Ctx) error // proxy işlemini yapar ve error döner
}

var Proxies = []Proxy{ // burada Proxy interface'ini implemente eden tüm struct'lar yer alır ve bu struct'lar ProxyHandler içinde dolaşılır
	NewLimitProxy("user", 3*time.Minute),
	//NewCacheProxy("user", 3*time.Minute),
	//NewCacheProxy("event", 3*time.Second),
}

func ProxyHandler(c *fiber.Ctx) error {
	for _, v := range Proxies { // Proxies içindeki tüm proxy'leri dolaş
		if v.Accept(c.Params("key")) { // eğer key bu proxy için geçerliyse
			return v.Proxy(c) // proxy işlemini yap ve error döner
		}
	}
	c.Response().SetStatusCode(404) // eğer hiçbir proxy için geçerli değilse 404 döner
	return nil
}
