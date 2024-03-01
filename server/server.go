package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	connections := make(map[int]net.Conn, 1024)
	i := 0

	listening, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Error: %s", err)
	}

	fmt.Println("Server start!!!")

	for {
		connect, err := listening.Accept()
		if err != nil {
			fmt.Print("Error: %s", err)
		}
		connections[i] = connect

		go exchange(connections, i)
		i++
	}
}

func exchange(connections map[int]net.Conn, n int) {
	connect := connections[n]

	var clientNo int

	buf := make([]byte, 256)

	fmt.Println("Accept connection:", n)
	defer connect.Close()

	for {
		read_len, err := connections[n].Read(buf)
		if err != nil {
			fmt.Println(err)
		}

		var message string
		message = string(buf[:read_len])

		_, err = fmt.Sscanf(message, "%d", &clientNo)
		if err != nil {
			connect.Write([]byte("error format message\n"))
			continue
		}
		pos := strings.Index(message, " ")
		if pos > 0 {
			out_message := message[pos+1:]

			connect = connections[clientNo]
			if connect == nil {
				connections[n].Write([]byte("client is close"))
				continue
			}

			out_buf := []byte(fmt.Sprintf("%d->>%s\n", clientNo, out_message))

			_, err := connect.Write(out_buf)
			if err != nil {
				fmt.Println("Error:", err.Error())
				break
			}
		}
	}
}
