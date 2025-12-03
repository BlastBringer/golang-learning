package main 

import (
	"fmt"
	"io"
	"net"
	"strings"
)


func main(){
	conn, err := net.Dial("tcp", ":8080")
	if err != nil{
		fmt.Println("not able to connect:" ,err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to the server!")
	data := "Hello, server!"
	if _, err := io.Copy(conn, strings.NewReader(data)); err != nil{
		fmt.Println("Error sending data:", err)
		return 
	}

	buf := make([]byte, len(data))
	if _, err := io.ReadFull(conn, buf); err != nil{
		fmt.Println("error reading data:", err)
		return 
	}

	fmt.Println("Received from server:", string(buf))

}

