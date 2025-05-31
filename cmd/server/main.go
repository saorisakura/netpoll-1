package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/netpoll"
)

type EchoHandler struct{}

// OnPrepare 准备处理连接
func (h *EchoHandler) OnPrepare(conn netpoll.Connection) context.Context {
	return context.Background()
}

// OnReadable 处理可读事件
func (h *EchoHandler) OnRequest(ctx context.Context, conn netpoll.Connection) error {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return err
	}
	if n == 0 {
		return nil
	}

	// 将接收到的数据原样返回给客户端
	_, err = conn.Write(buf[:n])
	return err
}

func NewEchoHandler() *EchoHandler {
	handler := new(EchoHandler)
	return handler
}

func main() {
	fmt.Println("hello netpoll.")
	listener, err := netpoll.CreateListener("tcp", "127.0.0.1:8080")
	if err != nil {
		panic("failed to create listener")
	}

	handler := NewEchoHandler()

	eventLoop, _ := netpoll.NewEventLoop(
		handler.OnRequest,
		netpoll.WithOnPrepare(handler.OnPrepare),
		netpoll.WithReadTimeout(time.Second),
	)

	eventLoop.Serve(listener)

	// stop server ...
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	eventLoop.Shutdown(ctx)
}
