package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	PORT := os.Getenv("PORT")

	listen, err := net.Listen("tcp", "0.0.0.0:"+PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()

	fmt.Println("listening on", PORT)
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go conn.Close()
	}
}
