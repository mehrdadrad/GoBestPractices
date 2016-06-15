package main

type geometry interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width, height float64
}

func (c circle) area() float64 {
	return 3.14159 * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func calcArea(g geometry) float64 {
	return g.area()
}

func main() {
	r := calcArea(circle{radius: 4})
	println(r)
	r = calcArea(rectangle{width: 2, height: 3})
	println(r)
}
