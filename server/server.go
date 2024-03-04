package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
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

	// Closes the connection.
	defer connect.Close()

	// Sending the ID to the user.
	num := []byte("Ваш ID: " + strconv.Itoa(n))
	_, err := connect.Write(num)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Подключисля юзер с ID: ", n)

	for {
		forWhom, fromWhom, message := breakingTheMessage(connections, n)
		connections[forWhom].Write([]byte(fmt.Sprintf("%d ->> %s", fromWhom, message)))
	}
}

func breakingTheMessage(connections map[int]net.Conn, n int) (forWhom, fromWhom int, b []byte) {
	// Reading the message.
	buf := make([]byte, 256)
	read_len, err := connections[n].Read(buf)
	if err != nil {
		fmt.Println(err)
	}

	// We remove all unnecessary things.
	var message string
	message = string(buf[:read_len])

	// Convert the received message, split it into recipient_message_sender.
	text := strings.Split(message, " ")

	forWhom, err = strconv.Atoi(text[0])
	if err != nil {
		log.Println("Не удалось преобразовать text[0] в forWhom", err)
	}
	fromWhom, err = strconv.Atoi(strings.Trim(text[len(text)-1], "\r"))
	if err != nil {
		log.Println("Не удалось преобразовать text[len(text)-1] в fromWhom", err)
	}

	mes := strings.TrimLeft(message, text[0])
	mes = strings.TrimRight(mes, text[len(text)-1])
	mes = strings.TrimSpace(mes)

	return forWhom, fromWhom, []byte(mes)
}
