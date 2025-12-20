package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {

	fmt.Println("Hello")
	fmt.Println(`This
			is
			a
			`)

	a := 42
	fmt.Println(a)

	b, c, d, _, f1 := 0, 1, 2, 3, "Happiness"
	fmt.Println(b, c, d, f1)

	var g int
	fmt.Println(g)

	adams := 42
	fmt.Printf("42 as binary, %b \n", adams)
	fmt.Printf("42 as hexadecimal, %x \n", adams)

	a, b, c, d = 0, 1, 2, 3
	e := 4
	f := 5

	fmt.Printf("%v \t %b \t %X \n", a, a, a)
	fmt.Printf("%v \t %b \t %X \n", b, b, b)
	fmt.Printf("%v \t %b \t %X \n", c, c, c)
	fmt.Printf("%v \t %b \t %X \n", d, d, d)
	fmt.Printf("%v \t %b \t %X \n", e, e, e)
	fmt.Printf("%v \t %b \t %X \n", f, f, f)

	y := 42
	z := 42.0
	fmt.Printf("%v of type %T \n", y, y)
	fmt.Printf("%v of type %T \n", z, z)

	var m float32 = 43.742
	fmt.Printf("%v of type %T \n", m, m)

	//z = m
	//fmt.Printf("%v of type %T \n", z, z)

	z = float64(m)
	fmt.Printf("%v of type %T \n", z, z)

	// Exercise 2

	num := rand.Intn(1000)
	fmt.Println("Number = ", num)

	sqr := math.Sqrt(2)
	fmt.Println("Square", sqr)

	pi := math.Pi
	fmt.Println("Pi", pi)

	fmt.Println("Addition functon:", Add(4, 3))

	fmt.Println("Multiplaction function", Multi(20, 5))

	a1, b1 := Swap("hello", "world")
	fmt.Println(a1, b1)

	x := 0
	x, y = Split(17)

	fmt.Println("Split ", x, y)

	fmt.Println("NeedInt ", NeedInt(Small))
	fmt.Println("needFloat small", NeedFloat(Small))
	fmt.Println("needFloat Big", NeedFloat(Big))

	fmt.Println(y) // Var for zero value

	z = 10 // Short declaration operator
	fmt.Println("Declaration ", z)

	a, b, c = 2, 4, 6 // Multiple initializations
	fmt.Println("Multiple", a, b, c)

	var y1 float64 = 100.01 // Var when specificity is required
	fmt.Println("specific", y1)

	_ = 20 // Blank indentifier

	a, b, c = 747, 911, 90210

	fmt.Printf("%d \t %b \t %#X \n", a, a, a)
	fmt.Printf("%d \t %b \t %#X \n", b, b, b)
	fmt.Printf("%d \t %b \t %#X \n", c, c, c)

	fmt.Println(Bark())

}
