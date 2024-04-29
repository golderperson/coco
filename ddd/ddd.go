package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	var d int
	var s string
	var f string

	_ = d

	// Listen for incoming connections on port 8080
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Accept incoming connections and handle them
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Print("検索語を入力してください: ")

		fmt.Scanf("%s", &s)
		d = len(s)
		f = strconv.Itoa(d)
		fmt.Println(f)
		_, err = conn.Write([]byte(f))

		_, err = conn.Write([]byte(s))
		fmt.Println(s)
		//here
		// Handle the connection in a new goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// Close the connection when we're done
	defer conn.Close()

	//listen
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the incoming data
	fmt.Printf("Received: %s", buf)
}
