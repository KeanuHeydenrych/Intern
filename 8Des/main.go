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

}
