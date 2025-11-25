package structs 

import "math"

type Rectangle struct{
	Width float64 
	Height float64
}

type Circle struct{
	Radius float64
}

func (c Circle) Area() float64{
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle)Area()float64{
	return r.Height * r.Width
}

