package main

// "strconv"

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
					// basically like anon structs.

					// Exercise:
					fmt.Println("\n\nExercise:")

					// Struct with slice
					p3 := person3{
						first: "James",
						last:  "Donoldson",
						age:   17,
						favIc: []string{"chocolate", "vanilla", "passion fruit with guava"},
					}

					fmt.Println(p3)

					for _, v := range p3.favIc {
						fmt.Println(p3.first, "'s fav icecream is:", v)
					}

					// map with struct
					m3 := map[string]person3{
						p1.last: p3,
						p2.last: p3,
					}

					for _, v := range m3 {
						fmt.Println(v)

						for _, v2 := range v.favIc {
							fmt.Println(v.first, v.last, v2)
						}
					}

					//embed sruct
					type engine struct {
						electric bool
					}

					type vehicle struct {
						engine
						make   string
						model  string
						doors  int
						colour string
					}

					v1 := vehicle{
						engine: engine{
							electric: true,
						},
						make:   "Ford",
						model:  "Mustang",
						doors:  4,
						colour: "Blue",
					}
					v2 := vehicle{
						engine: engine{
							electric: false,
						},
						make:   "BMW",
						model:  "X3-M",
						doors:  5,
						colour: "White",
					}

					fmt.Println(v1)
					fmt.Println(v2)

					fmt.Println(v1.electric, v1.make, v1.model)
					fmt.Println(v2.electric, v2.make, v2.model)

					//anon struct
					person1 := struct {
						first   string
						friends map[string]int
						drinks  []string
					}{
						first: "James",
						friends: map[string]int{
							"Goldfinger":    27,
							"W":             87,
							"Austin Powers": 52,
						},
						drinks: []string{
							"Gin - Tonic",
							"Brandy  - Coke",
						},
					}

					for k, v := range person1.friends {
						fmt.Println(person1.first, "- friends -", k, v)
					}

					for _, v := range person1.drinks {
						fmt.Println(person1.first, "- drinks -", v)
					}



				// Functons:
				fmt.Println("\n\nFunctions:")
				sum1 := sum(5, 6)
				fmt.Println("Sum: ", sum1)

				x2, y2 := 8, 999
				fmt.Println(x2, "+", y2, "=", sum(x2, y2))

				//Variadic Parameter
				sum_vari := sum_variadic(1, 2, 3, 4)
				fmt.Println(sum_vari)



			// Defer statement:		defer will cause the function to run last, like opening a file last
			defer foo() // to run this last
			bar()



		//10 Des
		// Methods:
		fmt.Println("\nMethods")

		p := person2{
			first: "James",
			last:  "Bond",
		}

		p.speak() // speak() is the method declared

		// Interfaces and polymorphism
		sa := secretAgent{
			person2: person2{
				first: "James",
				last:  "Bond",
				age:   007,
			},
			ltk: true,
		}

		sa.speak()

		saySomething(p) // because saySomething is of type 'Human', that contains 'speak()', any object with speak() can be called
		saySomething(sa)

		// stringer interface:
		b := book{
			title: " Nightwatch",
		}

		var n count = 42

		fmt.Println(b)
		fmt.Println(n)

		// Wrapper func for logging
		//			logInfo(b)
		//			logInfo(n)

		// Writer interface & writing to a file
		//													 Look at write_to_file.go!

		// anonymous func
		func() {
			fmt.Println("Anonymous func ran")
		}()

		func(s string) {
			fmt.Println("Anonymous Name:", s)
		}("Todd")

		// wrapper functions:
		fmt.Println("\nWrapper Function")

		add_a_b := doMath(8, 9, add)
		fmt.Println("a + b =", add_a_b)

		subtract_a_b := doMath(8, 9, subtract)
		fmt.Println("a - b =", subtract_a_b)

		// Pointers:
		fmt.Println("\nPointers:")

		x := 69
		fmt.Println("True Value:", x)
		fmt.Printf("Pointer: %v \t Type: %T", &x, &x)

		x_str := "Keanu"
		fmt.Println("True Value:", x_str)
		fmt.Printf("Poniter: %v \t Type: %T", &x_str, &x_str)

		// Dereferencing pointers:
		fmt.Println("\nDereferening pointers:")
		y := &x
		fmt.Printf("Pointer Value: %v \t Type of y: %T\n", y, y)
		fmt.Println("*y =", *y)
		//fmt.Println("*&x =", *&x)

		//Pass by value, pointers / ref types and mutability
		a := 70
		fmt.Print("\nOriginal ", a)
		intDelta(&a)
		fmt.Print(" intData ", a, "\n")

		xi := []int{1, 2, 3, 4}
		fmt.Print("Original ", xi)
		sliceData(xi)
		fmt.Print(" sliceData ", xi, "\n")

		m := make(map[string]int)
		m["James"] = 534
		fmt.Print("Original ", m["James"])
		mapDelta(m, " James ")
		fmt.Print(" mapDelta ", m["James"], "\n")

		// Value Semantics
		b1 := 1
		fmt.Println("Value Semantics")
		fmt.Println(addOne(b1))

		// Pointer Semantics
		fmt.Println("Pointer Semantics:")
		addOneP(&b1)
		fmt.Println("addOneP(&b1) has no value to print")

	*/

	// 11 Des

	// Value and Pointer Semantics

}

/*

type person2 struct {
	first string
	last  string
	age   int
}



type person3 struct {
	first string
	last  string
	age   int
	favIc []string
}


func sum(x int, y int) int {
	return x + y
}

// Variadic Parameter - takes on unlimited number of parameters
// func (r reciever) identifier(p paramters) (return(s)) {code}

func sum_variadic(ii ...int) int {
	fmt.Println("the sum of :", ii, "is =")

	n := 0
	for _, v := range ii {
		n += v
	}

	return n
}

// defer
func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}



type secretAgent struct {
	person2
	ltk bool
}

// methods
func (p person2) speak() {
	fmt.Println("I am", p.first)
}

func (sa secretAgent) speak() {
	fmt.Print("The name is ", sa.last, ". ", sa.first, " ", sa.last, ". Agent number: ")
	fmt.Printf("%03d\n", sa.age)
}

// Polymorphism
type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

// Stringer
type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("\nThe title is:", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("The number is:", strconv.Itoa(int(c))) //In terms of ASCII - helps convert int to string
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

// wrapper function
func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

// Pass by value
func intDelta(n *int) {
	*n = 73
}

func sliceData(ii []int) {
	ii[0] = 99
}

func mapDelta(md map[string]int, s string) {
	md[s] = 987
}

// Value Semantics
func addOne(v int) int {
	return v + 1
}

// Pointer Semantics
func addOneP(v *int) {
	*v += 1
}

*/
