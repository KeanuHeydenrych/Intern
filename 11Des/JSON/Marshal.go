package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// JSON documentation
	fmt.Println("\n\nJSON Documentation:")

	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   007,
	}

	p2 := person{
		First: "Gold",
		Last:  "Finger",
		Age:   69,
	}

	people := []person{p1, p2}

	fmt.Println("People:", people)
	// JSON: Marshal
	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	// JSON: Unmarshal
	s := `[{"First":"James","Bond","Age":007}, {"First":"Gold", "Last":"Finger","Age":69}]`
	bs = []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	people2 := []person2{}
	// var people2 []person2		another way to do this

	err = json.Unmarshal(bs, &people2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("all of the data", people2)

	for i, v := range people2 {
		fmt.Println("---- Person number:", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}

// JSON

type person struct {
	First string
	Last  string
	Age   int
}

// JSON unmarshal
type person2 struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}
