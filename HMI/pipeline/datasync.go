package pipeline

import (
	"encoding/hex"
	"fmt"
	"log"

	g "github.com/AllenDang/giu"
	"github.com/gorilla/websocket"
	"github.com/sammyoina/stewart-platform-ui/api"
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
			switch event := e.Event.(type) {
			case *models.SensorEvent_IMUEvent:
				fmt.Println("got data: ", event.IMUEvent.Pitch, event.IMUEvent.Yaw, event.IMUEvent.Roll)
			case *models.SensorEvent_PitotEvent:
				fmt.Println("got data: ", event.PitotEvent.DiffuserPitot, event.PitotEvent.IntakePitot, event.PitotEvent.TestSectionPitot)
			case *models.SensorEvent_StrainEvent:
				fmt.Println("got data: ", event.StrainEvent.Strain1, event.StrainEvent.Strain2, event.StrainEvent.Strain3, event.StrainEvent.Strain4, event.StrainEvent.Strain5, event.StrainEvent.Strain6)
			default:
				fmt.Println("no sensor event found")
				fmt.Println(hex.EncodeToString(message))
			}
		}
	}
}

type STDSender struct {
	conn *websocket.Conn
}

func (h *STDSender) StartOutputting(q queue.Queue) {
	fmt.Println("Start sending")
	for message, ok := q.Dequeue(); ok == true; message, ok = q.Dequeue() {

		h.conn = api.WebsocketConn
		if h.conn == nil {
			err := fmt.Errorf("Connection not established yet")
			g.Msgbox("Error", err.Error())
			break
		}
		err := h.conn.WriteMessage(websocket.BinaryMessage, message)
		if err != nil {
			fmt.Println("err: ", err)
			g.Msgbox("Error", err.Error())
			break
		}
		fmt.Println("Message sent: ", string(message), len(message))
	}
}
