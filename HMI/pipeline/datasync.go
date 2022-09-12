package pipeline

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/sammyoina/stewart-platform-ui/models"
	"github.com/sammyoina/stewart-platform-ui/queue"
	"google.golang.org/protobuf/proto"
)

type STDSync struct {
}

func (o *STDSync) StartOutputting(q queue.Queue) {
	fmt.Println("Starting output")
	for {
		for message, ok := q.Dequeue(); ok == true; message, ok = q.Dequeue() {
			var e models.SensorEvent
			if err := proto.Unmarshal(message, &e); err != nil {
				log.Println("failed to unmarshal:", err)
				return
			}
			fmt.Println("raw ", e.String())
			iMUev := e.GetIMUEvent()
			pitotev := e.GetPitotEvent()
			strainev := e.GetStrainEvent()
			if iMUev != nil {
				fmt.Println("got data: ", iMUev.Pitch, iMUev.Yaw, iMUev.Roll)
			}
			if pitotev != nil {
				fmt.Println("got data: ", pitotev.DiffuserPitot, pitotev.IntakePitot, pitotev.TestSectionPitot)
			}
			if strainev != nil {
				fmt.Println("got data: ", strainev.Strain1, strainev.Strain2, strainev.Strain3, strainev.Strain4, strainev.Strain5, strainev.Strain6)
			}
		}
	}
}

type STDSender struct {
	Conn *websocket.Conn
}

func (h *STDSender) StartOutputting(q queue.Queue) {
	fmt.Println("Start sending")
	for {
		for message, ok := q.Dequeue(); ok == true; message, ok = q.Dequeue() {
			if h.Conn == nil {
				fmt.Println("Conn not established yet")
				continue
			}
			err := h.Conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("Message sent")
		}
	}
}
