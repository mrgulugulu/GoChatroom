package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

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
	// wg.Add(2)
	// go read(conn)
	// go writeM(conn)
	// wg.Wait()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			return
		}
		fmt.Printf("received: %s\n", message)
	}
}

func timeWriter(conn *websocket.Conn) {
	for {
		time.Sleep(time.Second * 2)
		conn.WriteMessage(websocket.TextMessage, []byte(time.Now().Format("2006-01-02 15:04:05")))
	}
}

func writeM(conn *websocket.Conn) {
	defer wg.Done()
	for {
		fmt.Print("请输入:")
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
		fmt.Println("获取到的信息:", string(msg))
	}
}
