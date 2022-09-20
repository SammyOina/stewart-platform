package kinematics

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/sammyoina/stewart-platform-ui/pipeline"
	"github.com/sammyoina/stewart-platform-ui/queue"
)

func SetOrientation(yaw float64, pitch float64, roll float64, x float64, y float64, z float64) {
	yaw = d2r(yaw)
	pitch = d2r(pitch)
	roll = d2r(roll)
	plat := NewStewartPlatform(BASE_RADIUS, PLATFORM_RADIUS, d2r(HALF_ANGLE_BETWEEN_BASE), d2r(HALF_ANGLE_BETWEEN_PLATFORM), SERVO_HORN_LENGTH, ROD_LENGTH, 0)
	pos := plat.Calculate(yaw, roll, pitch, x, y, z)
	fmt.Println("pos", pos.Servo1)
	//pipeline.ServoPositionChannel <- pos
	//var wg sync.WaitGroup
	//pipeline.Wg = &wg
	//pipeline.Wg.Add(1)
	q := queue.NewChannelQueue()
	message, err := proto.Marshal(&pos)
	if err != nil {
		fmt.Println(err)
		return
	}
	q.Enqueue(message)
	sender := pipeline.STDSender{}
	sender.StartOutputting(q)
}
