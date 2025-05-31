package main

import (
	"fmt"
	"time"

	"github.com/cloudwego/netpoll"
)

func main() {
	fmt.Println("hello")
	client := netpoll.NewDialer()
	conn, err := client.DialConnection("tcp", "127.0.0.1:8080", time.Second*5)
	if err != nil {
		panic("failed to connect to server")
	} else {
		fmt.Println("connect success")
	}
	defer conn.Close()

	conn.Write([]byte("hello from client"))

	buf := make([]byte, 8192)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("read failed")
	} else {
		fmt.Println("read count ", n)
		fmt.Println(string(buf))
	}
}
