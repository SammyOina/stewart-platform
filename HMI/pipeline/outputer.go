package pipeline

import (
	"fmt"
	"log"
	"time"

	"github.com/sammyoina/stewart-platform-ui/models"
	"github.com/sammyoina/stewart-platform-ui/queue"
	"google.golang.org/protobuf/proto"
)

type StdoutOutputter struct {
}

func (o *StdoutOutputter) StartOutputting(q queue.Queue) {
	fmt.Println("Starting output")
	for {
		for message, ok := q.Dequeue(); ok == true; message, ok = q.Dequeue() {
			var e models.TestMessage
			if err := proto.Unmarshal(message, &e); err != nil {
				log.Println("failed to unmarshal:", err)
				return
			}
			fmt.Println("Got data: ", e.GetTestNumber())
		}
		time.Sleep(time.Second)
	}
}
