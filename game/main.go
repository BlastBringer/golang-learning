package main

import (
	"fmt"

)

func add(a, b int, ch chan int){
	ch <- a+b
}

func sender(ch chan int){
	ch <- 100
}

//lets learn and make some channels and goroutines
func main(){
	ch := make(chan int)
	ch2 := make(chan int)
	go add(2, 4, ch)
	go sender(ch2)

	select{
	case v := <-ch:
		fmt.Println(v)
	case v := <-ch2:
		fmt.Println(v)
	default:
		fmt.Print("nodata")
	}

	
}