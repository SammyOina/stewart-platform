package pipeline

import (
	"fmt"
	"io"
	"net"

	"github.com/sammyoina/stewart-platform-ui/queue"
)

type TcpListener struct {
	port int
}

func NewTcpListener(port int) *TcpListener {
	return &TcpListener{port: port}
}

func (l *TcpListener) StartAccepting(q queue.Queue) {
	fmt.Printf("Starting TCP listening on port %d\n", l.port)

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", l.port))
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go readFromConn(conn, q) // incl process buffer manually, retry logic, etc
	}
}

func readFromConn(c net.Conn, q queue.Queue) {
	var add = make(chan ([]byte), 1024)

	//go processBuffer(add, q)

	for {
		msg := make([]byte, 1024)
		rLen, err := c.Read(msg)

		if err != nil {
			if err == io.EOF {
				c.Close()
				return
			}

			panic(err)
		}

		add <- msg[:rLen]
		q.Enqueue(msg[:rLen])
	}
}
