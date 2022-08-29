package pipeline

import (
	"fmt"
	"time"

	"github.com/sammyoina/stewart-platform-ui/queue"
)

type StdoutOutputter struct {
}

func (o *StdoutOutputter) StartOutputting(q queue.Queue) {
	fmt.Println("Starting output")
	for {
		for message, ok := q.Dequeue(); ok == true; message, ok = q.Dequeue() {
			fmt.Println(fmt.Sprintf("Got data: %s", string(message)))
		}
		time.Sleep(time.Second)
	}
}
