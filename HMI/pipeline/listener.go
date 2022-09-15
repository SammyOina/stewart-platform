package pipeline

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sammyoina/stewart-platform-ui/api"
	"github.com/sammyoina/stewart-platform-ui/models"
	"github.com/sammyoina/stewart-platform-ui/queue"
	"google.golang.org/protobuf/proto"
)

type WebsocketListener struct {
	router *gin.Engine
	route  string
}

func NewWebsocketListener(router *gin.Engine, route string) *WebsocketListener {
	return &WebsocketListener{
		router: router,
		route:  route,
	}
}

func (l *WebsocketListener) StartAccepting(q queue.Queue) {

	apiHandler := api.MessageHandler{
		MessageQueue: q,
	}
	l.router.GET(l.route, apiHandler.DefaultHandler)
}

type StewartPositionListener struct {
	Pos chan models.ServoPositionEvent
}

func (s *StewartPositionListener) StartAccepting(q queue.Queue) {
	//defer Wg.Done()
	fmt.Println("start pos")
	var event models.ServoPositionEvent
	event = <-s.Pos
	message, err := proto.Marshal(&event)
	if err != nil {
		fmt.Println(err)
		return
	}
	q.Enqueue(message)
	fmt.Println("end pos")
	return
}
