package pipeline

import (
	"sync"

	"github.com/sammyoina/stewart-platform-ui/api"
	"github.com/sammyoina/stewart-platform-ui/models"
	"github.com/sammyoina/stewart-platform-ui/queue"
)

var ServoPositionChannel = make(chan models.ServoPositionEvent, 0)
var Wg sync.WaitGroup

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

	i := NewWebsocketListener(r, "/imu")
	q := queue.NewChannelQueue()
	o := &STDSync{}
	p := NewProcessor(i, q, o)
	go p.Start()

	i2 := &StewartPositionListener{
		Pos: ServoPositionChannel,
	}
	q2 := queue.NewChannelQueue()
	o2 := &STDSender{
		Conn: api.WebsocketConn,
	}
	p2 := NewProcessor(i2, q2, o2)
	go p2.Start()

	r.Run()
}

func NewProcessor(i inputter, q queue.Queue, o outputter) *Processor {
	var wg sync.WaitGroup
	wg.Add(1)
	return &Processor{i: i, q: q, o: o, wg: &wg}
}
