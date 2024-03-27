package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type Client struct {
	conn       net.Conn
	connected  bool
	serverIP   *walk.LineEdit
	serverPort *walk.LineEdit
	username   *walk.LineEdit
	mainWindow *walk.MainWindow
	chatTE     *walk.TextEdit
	sendTE     *walk.TextEdit
	connectBtn *walk.PushButton
	sendBtn    *walk.PushButton
	encryptCb  *walk.CheckBox
	keyEdit    *walk.LineEdit
}

func (c *Client) connectToServer() error {
	conn, err := net.Dial("tcp", c.serverIP.Text()+":"+c.serverPort.Text())
	if err != nil {
		return err
	}
	c.conn = conn
	c.connected = true
	c.updateUI()
	return nil
}

func (c *Client) disconnectFromServer() {
	if c.conn != nil {
		c.conn.Close()
		c.conn = nil
		c.connected = false
		c.updateUI()
	}
}

func (c *Client) sendMessage() {
	if c.connected {
		msg := c.sendTE.Text()
		if msg != "" {
			if c.encryptCb.Checked() {
				key, err := parseKey(c.keyEdit.Text())
				if err != nil {
					walk.MsgBox(c.mainWindow, "Ошибка", "Неверный ключ шифрования: "+err.Error(), walk.MsgBoxIconError)
					return
				}
				msg = permute(msg, key, true)
			}
			_, err := c.conn.Write([]byte(c.username.Text() + ": " + msg))
			if err != nil {
				log.Println("Ошибка отправки сообщения:", err)
				return
			}
			c.sendTE.SetText("")
		}
	}
}

func (c *Client) readLoop() {
	buf := make([]byte, 2048)
	for {
		n, err := c.conn.Read(buf)
		if err != nil {
			log.Println("Ошибка чтения:", err)
			c.disconnectFromServer()
			return
		}
		msg := string(buf[:n])
		if c.encryptCb.Checked() {
			key, err := parseKey(c.keyEdit.Text())
			if err != nil {
				walk.MsgBox(c.mainWindow, "Ошибка", "Неверный ключ шифрования: "+err.Error(), walk.MsgBoxIconError)
				return
			}
			msgParts := strings.SplitN(msg, ": ", 2)
			if len(msgParts) == 2 {
				msg = msgParts[0] + ": " + permute(msgParts[1], key, false)
			}
		}
		c.mainWindow.Synchronize(func() {
			c.chatTE.AppendText(msg + "\r\n")
		})
	}
}

func (c *Client) updateUI() {
	if c.connected {
		c.connectBtn.SetText("Отключиться")
		c.sendBtn.SetEnabled(true)
	} else {
		c.connectBtn.SetText("Подключиться")
		c.sendBtn.SetEnabled(false)
	}
}

func parseKey(keyStr string) ([]int, error) {
	keyStr = strings.TrimSpace(keyStr)
	key := make([]int, len(keyStr))
	seen := make(map[int]bool)

	for i, char := range keyStr {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			return nil, fmt.Errorf("Недопустимый ключ: %w", err)
		}
		if num < 1 || num > len(keyStr) {
			return nil, fmt.Errorf("Недопустимый ключ: число %d больше чем длина ключа", num)
		}
		if seen[num] {
			return nil, fmt.Errorf("Недопустимый ключ: число %d не уникальное", num)
		}
		seen[num] = true
		key[i] = num
	}

	return key, nil
}

func permute(text string, key []int, encrypt bool) string {
	keyLength := len(key)

	for len(text)%keyLength != 0 {
		text += " "
	}

	result := make([]byte, len(text))

	for i := 0; i < len(text); i += keyLength {
		block := []byte(text[i : i+keyLength])
		permutedBlock := make([]byte, keyLength)

		if encrypt {
			for j, k := range key {
				permutedBlock[k-1] = block[j]
			}
		} else {
			for j, k := range key {
				permutedBlock[j] = block[k-1]
			}
		}

		copy(result[i:], permutedBlock)
	}

	resultText := strings.TrimRight(string(result), " ")

	return resultText
}

func main() {
	client := &Client{}

	MainWindow{
		AssignTo: &client.mainWindow,
		Title:    "Чат",
		Size:     Size{Width: 600, Height: 400},
		Layout:   VBox{},
		Children: []Widget{
			Composite{
				Layout: HBox{},
				Children: []Widget{
					LineEdit{AssignTo: &client.serverIP, Text: "IP"},
					LineEdit{AssignTo: &client.serverPort, Text: "Port"},
					LineEdit{AssignTo: &client.username, Text: "UserName"},
					PushButton{
						AssignTo: &client.connectBtn,
						Text:     "Подключиться",
						OnClicked: func() {
							if !client.connected {
								err := client.connectToServer()
								if err != nil {
									walk.MsgBox(client.mainWindow, "Ошибка", "Не удалось подключиться к серверу", walk.MsgBoxIconError)
									return
								}
								go client.readLoop()
							} else {
								client.disconnectFromServer()
							}
						},
					},
					CheckBox{
						AssignTo: &client.encryptCb,
						Text:     "Шифровать сообщения",
					},
					LineEdit{
						AssignTo: &client.keyEdit,
						Text:     "Ключ",
					},
				},
			},
			TextEdit{AssignTo: &client.chatTE, ReadOnly: true},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					TextEdit{AssignTo: &client.sendTE},
					PushButton{
						AssignTo:  &client.sendBtn,
						Text:      "Отправить",
						OnClicked: client.sendMessage,
					},
				},
			},
		},
	}.Run()
}
