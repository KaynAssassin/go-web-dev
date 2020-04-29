package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main(){
	li,err := net.Listen("tcp",":8001")
	if err != nil{
		log.Fatal(err)
	}

	defer li.Close()
	for {
		conn,err := li.Accept()
		if err != nil{
			log.Println(err)
			continue
		}
		io.WriteString(conn,"\nHello from TCP server \n")
		fmt.Println(conn,"How is your today")
		fmt.Fprintf(conn,"%v","Well,i hoped!")

		conn.Close()
	}

}