package main

/*
import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main2() {

	fmt.Println("YEYE")

	p := person2{
		first: "James",
		last:  "Bond",
	}

	// Writer interface & writing to a file
	f, err := os.Create("temp.txt")
	if err != nil {
		log.Fatalf("error%s", err)
	}
	defer f.Close() // have to close a file if you open it!!!!

	s := []byte("Hello!")

	_, err = f.Write(s)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	// byte buffer
	byt := bytes.NewBufferString("\nThis is a byte")

	fmt.Println(byt.String())

	byt.Reset()

	fmt.Println(byt) // byt is empty

	var b5 bytes.Buffer
	p.writeOut(&b5)
}

// Wrapper for logging
func logInfo(s fmt.Stringer) {
	log.Println("LOG FROM 138", s.String())
}

// Writer interface
func (p person2) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(p.first))
	return err
}

*/
