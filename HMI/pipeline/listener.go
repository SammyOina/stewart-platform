package pipeline

import (
	"github.com/gin-gonic/gin"
	"github.com/sammyoina/stewart-platform-ui/api"
	"github.com/sammyoina/stewart-platform-ui/queue"
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
