package pipeline

import (
	"sync"

	"github.com/sammyoina/stewart-platform-ui/api"
	"github.com/sammyoina/stewart-platform-ui/queue"
)

type Queue interface {
	Enqueue(data []byte)
	Dequeue() ([]byte, bool)
}

type inputter interface {
	StartAccepting(q queue.Queue)
}

type outputter interface {
	StartOutputting(q queue.Queue)
}

type Processor struct {
	i  inputter
	q  queue.Queue
	o  outputter
	wg *sync.WaitGroup
}

func (p *Processor) Start() {
	go p.i.StartAccepting(p.q)
	go p.o.StartOutputting(p.q)
	p.wg.Wait()
}

func InitPipeline() {
	r := api.GetRouter()

	i3 := NewWebsocketListener(r, "/imu")
	q3 := queue.NewChannelQueue()
	o3 := &STDSync{}
	p3 := NewProcessor(i3, q3, o3)
	go p3.Start()

	r.Run()
}

func NewProcessor(i inputter, q queue.Queue, o outputter) *Processor {
	var wg sync.WaitGroup
	wg.Add(1)
	return &Processor{i: i, q: q, o: o, wg: &wg}
}
