package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		bt := make([]byte, 1024)
		_, err = conn.Read(bt)
		if err != nil {
			log.Fatal(err)
		}
		msg := string(bt)

		go func(msg string) {
			defer conn.Close()
			time.Sleep(time.Second)
			fmt.Println(conn.RemoteAddr().String(), "-->", msg)
			// resp := msg + " done"
			_, err = conn.Write([]byte(msg + " done"))
			if err != nil {
				log.Fatal(err)
			}
		}(msg)
	}
}
