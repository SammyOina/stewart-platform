package pipeline

import (
	"fmt"
	"net"

	"github.com/sammyoina/stewart-platform-ui/queue"
)

type UdpListener struct {
	port int
}

func NewUdpListener(port int) *UdpListener {
	return &UdpListener{port: port}
}

func (l *UdpListener) StartAccepting(q queue.Queue) {
	fmt.Printf("Starting UDP listening on port %d\n", l.port)

	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", l.port))
	if err != nil {
		fmt.Println(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	buf := make([]byte, 65536)

	for {
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error: ", err)
		}

		message := buf[0:n]
		q.Enqueue(message)
	}
}
