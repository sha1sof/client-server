package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

type Client struct {
	conn  net.Conn
	login string
}

func main() {
	clients := make(map[string]*Client)
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Ошибка при прослушивании: ", err)
	}
	defer listener.Close()

	fmt.Println("Сервер запущен!!!")

	var wg sync.WaitGroup

	for {
		connect, err := listener.Accept()
		if err != nil {
			log.Println("Ошибка при принятии подключения: ", err)
			continue
		}

		wg.Add(1)
		go handleConnection(connect, clients, &wg)
	}
}

func handleConnection(connect net.Conn, clients map[string]*Client, wg *sync.WaitGroup) {
	defer connect.Close()
	defer wg.Done()

	buf := make([]byte, 256)
	readLen, err := connect.Read(buf)
	if err != nil {
		log.Println("Ошибка при чтении логина: ", err)
		return
	}

	login := strings.TrimSpace(string(buf[:readLen]))

	if _, exists := clients[login]; exists {
		connect.Write([]byte("Логин уже занят\n"))
		return
	}

	client := &Client{conn: connect, login: login}
	clients[login] = client

	connect.Write([]byte("Подключение выполнено. Ваш логин: " + login + "\n"))

	fmt.Println("Новый клиент подключен. Логин:", login)

	for {
		readLen, err := connect.Read(buf)
		if err != nil {
			log.Println("Ошибка при чтении сообщения: ", err)
			delete(clients, login)
			break
		}

		message := string(buf[:readLen])
		fmt.Println("Получено сообщение от", login+":", message)

		parts := strings.Fields(message)
		if len(parts) < 2 {
			fmt.Println("Неправильный формат сообщения")
			continue
		}

		receiverLogin := parts[0]
		messageToSend := strings.Join(parts[1:], " ")

		receiverClient, exists := clients[receiverLogin]
		if !exists {
			fmt.Println("Пользователь с логином", receiverLogin, "не найден")
			continue
		}

		_, err = receiverClient.conn.Write([]byte(fmt.Sprintf("Сообщение от %s: %s\n", login, messageToSend)))
		if err != nil {
			fmt.Println("Ошибка при отправке сообщения клиенту", receiverLogin, ":", err)
		}
	}
}
