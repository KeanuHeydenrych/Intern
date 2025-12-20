package main

import (
	"fmt"
	"math/rand"
)

func main() {

	//Sequential control flow (statements):
	fmt.Println("\nFirst statement")
	fmt.Println("Second statement")

	r := rand.New(rand.NewSource(99))

	x := r.Int()
	y := 5
	fmt.Printf("First expression x = %v \nSecond expression y = %v\n", x, y)

	// Conditional control flow (if, switch loops)
	fmt.Println("\n\nConditional control flow:")

	fmt.Println("\nIf statements:")
	if x < 42 {
		fmt.Println("Less than 42")
	} else {
		fmt.Println("greater than 42")
	}

	fmt.Println("\nSwitch statements:")
	switch {
	case y < 5:
		fmt.Println("y is smaller than 5")
	case y > 5:
		fmt.Println("y is greater than 5")
	case y == 5:
		fmt.Println("y IS 5")
	}

	fmt.Println("\n\nIterative loops")
	// Iterative loops (only explicit FOR loop exists)
	// Do loops, While loops do not exist - use FOR loop to "DIY"

	fmt.Println("For loop i")
	for i := 0; i < 5; i++ {
		fmt.Println("i = ", i)
	}

	// 		fmt.Println(i) // Note that i is not declared outside the for loop

	//Other way to do for loop:
	fmt.Println("For loop j:")
	j := 6
	for j = range 10 {
		fmt.Println("j = ", j)
	}
	// note: The output does not start at j = 6, this is because the "j = range" starts the loop at j = 0 and does 10 iterations.
	// Use previous way if you want specific starting point

	// third way:
	fmt.Println("k outside for loop")
	k := 1
	for k <= 5 {
		fmt.Println("k = ", k)
		k++
	}

	fmt.Println("Final k = ", k) // Note that final k = 6, bc the final iteration increased k by one, after k <= 5

	/*
		fmt.Println("\n\nGo routine channels:")
		x_y_ch := make(chan int)
		y_x_ch := make(chan int)

		go func() {
			x_y_ch <- x - y
		}()

		go func() {
			y_x_ch <- y - x
		}()

		select {
		case s1 := <-x_y_ch:
			fmt.Println("x - y = ", s1)
		case s2 := <-y_x_ch:
			fmt.Println("y - x = ", s2)
		}
	*/

	//Aggregation:
	fmt.Println("\nAggregation:")

	//arrays
	fmt.Println("Arrays:")

	var ni [7]int //declare a variable of type [7]int

	ni[0] = 42 //assign a value to position 0 in index
	fmt.Printf("%#v \t\t\t\t %T\n", ni, ni)

	ni2 := [4]int{55, 56, 57, 58} //array literal
	fmt.Printf("%#v \t\t\t\t\t %T\n", ni2, ni2)

	fmt.Println(len(ni))
	fmt.Println(len(ni2))

	ns := [...]string{"chocolate", "strawberry", "vanilla"}
	fmt.Printf("%#v \t %T\n", ns, ns)

	fmt.Println(len(ns))

	//slices
	fmt.Println("\nSlices")
	xs := []string{"This", "is", "a", "slice"} //all have the same type (String)
	fmt.Printf("%#v \t\t\t %T\n ", xs, xs)

	for i := range xs {
		fmt.Println(xs[i])
	}

	xs = append(xs, "This is an", "apppend.")
	fmt.Println(xs)

	//making a slice with 'make'
	xi := make([]int, 0, 10)
	fmt.Println("Slice xi:", xi)

	fmt.Println("Length of xi", len(xi))
	fmt.Println("Cap of xi", cap(xi))

	xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println("appended xi:", xi)

	xsl := [5]int{}

	for i := 0; i < 5; i++ {
		xsl[i] = i + 10
	}

	for i, v := range xsl {
		fmt.Printf("%v \t %T \t %v\n ", v, v, i)
	}

	//maps
	fmt.Println("\nMaps")

	//structs
	fmt.Println("\nStructs")

}
