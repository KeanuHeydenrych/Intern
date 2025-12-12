package main

import (
	"fmt"
	"time"

	saying "girhub.com/KeanuHeydenrych/Intern/12Des/Benchmark"
)

func main() {

	fmt.Println("12Des")
	start := time.Now()

	/*

		// errors

		//var ans1, ans2, ans3 string
		answers := make([]string, 3)

		fmt.Println("Name:")
		_, err := fmt.Scan(&answers[0])
		if err != nil {
			fmt.Println("Error:", err)
		}

		fmt.Println("Last Name:")
		_, err = fmt.Scan(&answers[1])
		if err != nil {
			panic(err)
		}

		fmt.Println("Job")
		_, err = fmt.Scan(&answers[2])
		if err != nil {
			panic(err)
		}
		fmt.Println("\nInfo gained:")
		fmt.Println("Name:", answers[0])
		fmt.Println("Last Name:", answers[1])
		fmt.Println("Job:", answers[2])

	*/

	// Documentation:
	// godoc.org

	// Testing & Benchmarking

	fmt.Println("\n\nTesting & Benchmarking:")

	a := mySum(2, 3)
	b := mySum(2, 3, 6, 8)

	fmt.Println("2 + 3 =", a)
	fmt.Println("2 + 3 + 6 + 8 =", b)

	a2 := mySum2(2, 3)
	fmt.Println("2 + 3 =", a2) // this gives a testing error

	a3 := mySum3(1, 58)
	fmt.Println("1 + 58 =", a3) // this gives a testing error

	// Benchmarking:
	st := "Keanu"
	fmt.Println(saying.Greetings(st))

	// git bash timer
	elapsed := time.Since(start)
	fmt.Println("\n\nExecution time:", elapsed)

}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}

	return sum
}

func mySum2(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}

	return sum + 1
}

func mySum3(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}

	return sum + 1
}
