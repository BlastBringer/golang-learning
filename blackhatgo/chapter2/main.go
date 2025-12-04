package main 


import (
	"fmt"
	"time"
)

func main(){

	money := 0
	for{
		time.Sleep(time.Duration(time.Second) * 1)
		money++
		fmt.Println(makemoney(money))
	}
}


func makemoney(money int) string {
	return fmt.Sprintf("%d Rupees made on %d:%d hrs:mins", money, time.Now().Local().Hour(), time.Now().Local().Minute())
	
}