package main 
import (
	"fmt"
	"net"
)

func handleClient(conn net.Conn){
	defer conn.Close() 
	buffer := make([]byte, 1024)


	for {
		//read data from the clinet 
		n, err := conn.Read(buffer)
		if err != nil{
			fmt.Println("Error:", err)
			return 
		}
			fmt.Printf("Received : %s\n", buffer[:n])
	}

}