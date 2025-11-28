package main 

import (
	"fmt"
	"os"
	"log"
)

type FooReader struct{} 

func(fooReader *FooReader) Read(b []byte) (int, error){
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

type FooWriter struct{}

func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out> ")
	return os.Stdout.Write(b)
}

func main(){
	var (
		reader FooReader 
		writer FooWriter 
	)

	input := make([]byte, 4096)

	s, err := reader.Read(input)
	if err != nil{
		log.Fatalln("Unable to read data")
	}

	fmt.Printf("Read %d bytes from stdin\n", s)
	
	//use writer to write output
	s, err = writer.Write(input)
	if err != nil{
		log.Fatalln("Unable to write the data")
	}

	fmt.Printf("wrote %d bytes to stdout\n", s)
}