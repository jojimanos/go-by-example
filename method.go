package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	height, width float64
}

type circle struct {
	radius float64
}

func (r rectangle) Area() float64 {
	return r.height * r.width
}

func (c circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	c1 := circle{20}

	fmt.Println("This is the area", c1.Area())
}
