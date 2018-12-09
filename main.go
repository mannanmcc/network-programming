package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:4242")
	if err != nil {
		log.Fatal(err)
	}

	for {
		//take a new connection from client
		conn, err := lis.Accept()
		if err != nil {
			log.Fatal(err)
		}
		//let's print whatever is sent from client using io.copy
		go handleClientRequest(conn)
	}

}

func handleClientRequest(conn net.Conn) {
	n, err := io.Copy(os.Stderr, conn)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Copied %d bytes", n)
}
