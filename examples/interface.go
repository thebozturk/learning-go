package main

import "fmt"

// interface bize yapılmasını istediğimiz işlemleri söyler

type rectangle struct {
	a, b float64
}

func (r rectangle) area() float64 { // structa bir method tanımladık
	return r.a * r.b
}

func (r rectangle) circumference() float64 {
	return 2 * (r.a + r.b)
}

type shape interface {
	area() float64
	circumference() float64
} // interface sadece methodlarla ilgileniyor

type circle struct {
	r float64
}

func (c circle) area() float64 {
	return 3.14 * c.r * c.r
}

func (c circle) circumference() float64 {
	return 2 * 3.14 * c.r
}

func interfaceFunc(s shape) {
	fmt.Println("Area:", s.area())
	fmt.Println("Circumference:", s.circumference())
}

func main() {
	r1 := rectangle{a: 2, b: 3}
	interfaceFunc(r1)
	fmt.Printf("Type of r1: %T\n", r1)
	c1 := circle{r: 5}
	interfaceFunc(c1)
	fmt.Printf("Type of c1: %T\n", c1)
}
