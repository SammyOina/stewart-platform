package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sammyoina/stewart-platform-ui/queue"
)

type MessageHandler struct {
	MessageQueue queue.Queue
}

var WebsocketConn *websocket.Conn

var upgrader = websocket.Upgrader{}

func (m *MessageHandler) DefaultHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	WebsocketConn = conn
	if err != nil {
		return
	}

	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		//conn.WriteMessage(websocket.TextMessage, []byte("heyaa"))
		/*var e models.MessMan
		if err := proto.Unmarshal(message, &e); err != nil {
			log.Println("failed to unmarshal:", err)
			return
		}
		fmt.Println("Got data: ", e.Hey)*/
		if err != nil {
			fmt.Println("err ", err)

			break
		}
		m.MessageQueue.Enqueue(message)
	}
}
