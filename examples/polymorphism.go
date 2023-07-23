package main

import "fmt"

// polymorphism: çok biçimlilik demektir yani bir nesnenin birden fazla biçimi olmasıdır.
// sağladığı avantajlar: kod tekrarını önler, kodun okunabilirliğini arttırır, esneklik sağlar. (örneğin bir nesnenin birden fazla biçimi olduğu için o nesneye yeni bir biçim eklemek istediğimizde sadece yeni biçimi eklememiz yeterli olacaktır.)

type Shape interface {
	Area() float64
}

type Square struct {
	Length float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// gelen parametrelerin hepsi Shape interface'ini implemente ettiği için bu fonksiyonu kullanabiliriz.
func printArea(shapes ...Shape) {
	for _, shape := range shapes {
		fmt.Println("Alan:", shape.Area())
	}
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	t := Triangle{Base: 10, Height: 5}
	s := Square{Length: 10}
	r := Rectangle{Width: 10, Height: 5}

	printArea(t, s, r)
}
