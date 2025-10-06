package main

import ("fmt"
        "math"  )

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	height, breadth float64
}

func (v Circle) Area()float64 {
	return math.Pi*v.radius*v.radius
}

func (v Circle) Perimeter()float64 {
	return 2*math.Pi*v.radius
}

func (v Rectangle) Area()float64 {
	return v.height*v.breadth

}

func (v Rectangle) Perimeter()float64 {
	return 2*(v.height+v.breadth)

}

func Calculate(c Shape) {
   fmt.Println("area is ",c.Area(),"\n")
   fmt.Println("perimeter is ",c.Perimeter(),"\n")
}

func main() {
   c:=Circle{
	radius:100,
   }

   r:=Rectangle{
	height:100,
	breadth:200,
   }
   
   fmt.Println("Circle \n")
   Calculate(c)
    fmt.Println("Rectangle \n")
   Calculate(r)

}
