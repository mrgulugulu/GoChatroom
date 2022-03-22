package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
	"sync"

	"github.com/gorilla/websocket"
)

type Message struct {
	//消息struct
	Sender    string `json:"sender,omitempty"`    //发送者
	Recipient string `json:"recipient,omitempty"` //接收者
	Content   string `json:"content,omitempty"`   //内容
	ServerIP  string `json:"serverIp,omitempty"`
	SenderIP  string `json:"senderIp,omitempty"`
}

//定义连接的服务端的网址
var addr = flag.String("addr", "localhost:8080", "http service address")
var wg sync.WaitGroup

func main() {
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	var dialer *websocket.Dialer

	//通过Dialer连接websocket服务器
	conn, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// go timeWriter(conn)
	wg.Add(2)
	go read(conn)
	go writeM(conn)
	wg.Wait()

	// for {
	// 	_, message, err := conn.ReadMessage()
	// 	if err != nil {
	// 		fmt.Println("read:", err)
	// 		return
	// 	}
	// 	fmt.Printf("received: %s\n", message)
	// }
}

func writeM(conn *websocket.Conn) {
	defer wg.Done()
	for {

		fmt.Println("输入信息:")
		reader := bufio.NewReader(os.Stdin)
		data, _ := reader.ReadString('\n')
		conn.WriteMessage(1, []byte(data))
	}
}
func read(conn *websocket.Conn) {
	defer wg.Done()
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("错误信息:", err)
			break
		}
		if err == io.EOF {
			continue
		}
		message := &Message{}
		_ = json.Unmarshal(msg, message)
		if message.Recipient == "" {
			fmt.Printf("获取到广播信息, sender:%s, content:%s\n", message.Sender, message.Content)
		} else {
			fmt.Printf("获取来自用户%s的私信, content:%s\n", message.Sender, message.Content)
		}

	}
}
