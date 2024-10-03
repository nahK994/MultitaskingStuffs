package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Write([]byte("task 1"))
	if err != nil {
		log.Fatal(err)
	}

	bt := make([]byte, 2048)
	_, err = conn.Read(bt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bt))
}
