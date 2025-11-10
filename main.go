package main 

import (
	"time"
	"fmt"
	"math/rand"
)

func main(){
	ch := make(chan int)

	for i := 0; i < 3; i++{
		go func(id int) {
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			ch <- id
		}(i)
	}

	winner := <-ch 
	fmt.Println("winner is", winner)
}