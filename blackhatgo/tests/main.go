package main 

import(
	"fmt"
)

const names = {"Rahul"}

func Hello(name string) string{
	return	fmt.Sprintf("Hello, %s", name)
}

func main(){
	fmt.Println(Hello("rahul"))
	go Hello("rahul")
}