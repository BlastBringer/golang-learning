package main 

import (
	"fmt"
	"bufio"
	"net"
)


func main(){
	listener, err := net.Listen("tcp", ":9000")
	if err != nil{
		panic(err)
	}


	defer listener.Close()

	fmt.Println("Server listening on port 9000")

	for {
		conn, err := listener.Accept() 
		if err != nil{
			fmt.Println("Error accepting:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn){
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil{
			fmt.Println("Client Disconnected")
			return 
		}

		fmt.Printf("Received:%s", message)
		conn.Write([]byte("You said: "+ message))
	}
}