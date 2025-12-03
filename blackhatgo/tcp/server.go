package main 

import (
	"fmt"
	"net"
	"io"
)


func main(){
	listener , err := net.Listen("tcp", ":8080")
	if err != nil{
		fmt.Println("Error creating listener:", err)
		return 
	}

	defer listener.Close()

	fmt.Println("Listening on :8080....")

	for {
		conn, err := listener.Accept() 
		if err != nil{
			fmt.Println("Error accepting connection:", err)
			continue 
		}

		go handleConnection(conn)

	}
}


func handleConnection(conn net.Conn){
	defer conn.Close()
	fmt.Println("New connecton from", conn.RemoteAddr())

	//echo the data received back to the client
	if _, err := io.Copy(conn, conn); err != nil{
		fmt.Println("Error echoing data:", err)
	}
}