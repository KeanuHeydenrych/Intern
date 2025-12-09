package main

import (
	"fmt"
)

func main() {

	/*


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

	/*

		//Aggregation (array, slice, map, struct):
		fmt.Println("\n\nAggregation:")

		//arrays
		fmt.Println("Arrays:")

		var ni [7]int //declare a variable of type [7]int

		ni[0] = 42 //assign a value to position 0 in index
		fmt.Printf("ni: %#v \t\t\t\t %T\n", ni, ni)

		ni2 := [4]int{55, 56, 57, 58} //array literal
		fmt.Printf("ni2: %#v \t\t\t\t\t %T\n", ni2, ni2)

		fmt.Println("length of ni", len(ni))
		fmt.Println("length of ni2", len(ni2))

		ns := [...]string{"chocolate", "strawberry", "vanilla"}
		fmt.Printf("ns: %#v \t %T\n", ns, ns)

		fmt.Println("length of ns", len(ns))

		//slices
		fmt.Println("\nSlices:")
		xs := []string{"This", "is", "a", "slice"} //all have the same type (String)
		fmt.Printf("xs: %#v \t\t\t %T\n ", xs, xs)

		for i := range xs {
			fmt.Println("xs[i]:", xs[i])
		}

		xs = append(xs, "This is an", "apppend.")
		fmt.Println("xs", xs)

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
			fmt.Printf("xsl: %v \t %T \t %v\n ", v, v, i)
		}

	*/

	//maps
	fmt.Println("\nMaps:")
	m := map[string]int{ //[string] is used for the lookup value. int is the return value
		"Todd":  42,
		"Henry": 16,
		"Keanu": 23,
	}

	fmt.Println("The age of Keanu is:", m["Keanu"])

	fmt.Println("\nThe age of", m)

	m["Georg"] = 73 // append a map

	fmt.Println("Appended map:", m)

	fmt.Println("All var's in map:")
	for i, v := range m {
		fmt.Println(i, v)
	}

	fmt.Println("Follows order 2, 3, 4, 1?")
	for _, v := range m {
		fmt.Print(v, " ")
	}

	fmt.Println("\nOnly return values, i.e. age (int):")
	for k := range m {
		fmt.Println(m[k])
	}

	delete(m, "Georg") //delete an element form an map
	fmt.Println("/nDeleted one element", m)

	//looking up a value that does not exist
	v, ok := m["heyderych"]
	if ok {
		fmt.Println("the value exists", v)
	} else {
		fmt.Println("the value does not exist")
	}

	//structs
	fmt.Println("\nStructs:")

	// can also be outside func main() scope
	type person struct {
		first string
		last  string
		age   int
	}

	p1 := person{
		first: "James",
		last:  "Bond",
		age:   007,
	}

	p2 := person2{
		first: "Johnny",
		last:  "Bravo",
		age:   27,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Printf("%#v \t %T\n", p1, p1)
	fmt.Printf("%#v \t %T\n", p1, p1)

	//Embedded Structs
	type secretAgent struct {
		person
		ltk bool
	}

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   -007,
		},
		ltk: true,
	}

	fmt.Println(sa1)
	fmt.Println("Only first name", sa1.first)
	fmt.Printf("%#v \t %T\n", sa1, sa1)

	//Anonymous struct, useful for when you dont want to create another type
	c1 := struct {
		model string
		brand string
		year  int
	}{
		model: "R8",
		brand: "Audi",
		year:  2021,
	}

	fmt.Println("\nAnon Struct:", c1)
	fmt.Printf("%#v \t %T\n", c1, c1)

	//Composition

}

type person2 struct {
	first string
	last  string
	age   int
}
