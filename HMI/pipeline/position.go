package pipeline

import (
	"fmt"
	"log"

	g "github.com/AllenDang/giu"
	"github.com/golang/protobuf/proto"
	"github.com/sammyoina/stewart-platform-ui/kinematics"
	"github.com/sammyoina/stewart-platform-ui/queue"
)

const (
	ROD_LENGTH                  float64 = 14.0
	BASE_RADIUS                 float64 = 15.0
	PLATFORM_RADIUS             float64 = 10.0
	SERVO_HORN_LENGTH           float64 = 4.5
	HALF_ANGLE_BETWEEN_BASE     float64 = 13
	HALF_ANGLE_BETWEEN_PLATFORM float64 = 13
)

var ErrorMessage string

var Platform kinematics.StewartPlatform = kinematics.NewStewartPlatform(BASE_RADIUS, PLATFORM_RADIUS, kinematics.D2r(HALF_ANGLE_BETWEEN_BASE), kinematics.D2r(HALF_ANGLE_BETWEEN_PLATFORM), SERVO_HORN_LENGTH, ROD_LENGTH, 0)

func SetOrientation(yaw float64, pitch float64, roll float64, x float64, y float64, z float64) {
	yaw = kinematics.D2r(yaw)
	pitch = kinematics.D2r(pitch)
	roll = kinematics.D2r(roll)
	//plat := NewStewartPlatform(BASE_RADIUS, PLATFORM_RADIUS, d2r(HALF_ANGLE_BETWEEN_BASE), d2r(HALF_ANGLE_BETWEEN_PLATFORM), SERVO_HORN_LENGTH, ROD_LENGTH, 0)
	pos, err := Platform.Calculate(yaw, roll, pitch, x, y, z)
	if err != nil {
		log.Println(err)
		g.Msgbox("Error", err.Error())
		return
	}
	fmt.Println("pos", pos.Servo1, pos.Servo2, pos.Servo3, pos.Servo4, pos.Servo5, pos.Servo6)
	//pipeline.ServoPositionChannel <- pos
	//var wg sync.WaitGroup
	//pipeline.Wg = &wg
	//pipeline.Wg.Add(1)

	q := queue.NewChannelQueue()
	message, err := proto.Marshal(pos)
	if err != nil {
		log.Println(err)
		g.Msgbox("Error", err.Error())
		return
	}
	q.Enqueue(message)
	sender := STDSender{}
	sender.StartOutputting(q)
}
