package pipeline

import (
	"log"
	"time"

	g "github.com/AllenDang/giu"
	"github.com/golang/protobuf/proto"
	"github.com/sammyoina/stewart-platform-ui/kinematics"
	"github.com/sammyoina/stewart-platform-ui/queue"
)

func TestPlatformRoll() {
	q := queue.NewChannelQueue()
	plat := kinematics.NewStewartPlatform(BASE_RADIUS, PLATFORM_RADIUS, HALF_ANGLE_BETWEEN_BASE, HALF_ANGLE_BETWEEN_PLATFORM, SERVO_HORN_LENGTH, ROD_LENGTH, 0)

	pos, err := plat.Calculate(0, kinematics.D2r(float64(34)), 0, 0, 0, 0)
	if err != nil {
		log.Println(err)
		g.Msgbox("Error", err.Error())
	}
	message, err := proto.Marshal(pos)
	if err != nil {
		log.Println(err)
		g.Msgbox("Error", err.Error())
	}
	q.Enqueue(message)
	sender := STDSender{}
	sender.StartOutputting(q)
	time.Sleep(time.Millisecond * 1000)
	pos, err = plat.Calculate(0, kinematics.D2r(float64(-34)), 0, 0, 0, 0)
	if err != nil {
		log.Println(err)
		g.Msgbox("Error", err.Error())

	}
	message, err = proto.Marshal(pos)
	if err != nil {
		log.Println(err)
		g.Msgbox("Error", err.Error())
	}
	q.Enqueue(message)
	sender.StartOutputting(q)
	time.Sleep(time.Millisecond * 1000)
	g.Msgbox("Tests", "Roll test done")
}
func TestPlatformYaw() {
	q := queue.NewChannelQueue()
	plat := kinematics.NewStewartPlatform(BASE_RADIUS, PLATFORM_RADIUS, HALF_ANGLE_BETWEEN_BASE, HALF_ANGLE_BETWEEN_PLATFORM, SERVO_HORN_LENGTH, ROD_LENGTH, 0)

	pos, err := plat.Calculate(kinematics.D2r(float64(20)), 0, 0, 0, 0, 0)
	if err != nil {
		log.Println(err)
		g.Msgbox("Error", err.Error())
	}
	message, err := proto.Marshal(pos)
	if err != nil {
		log.Println(err)
		g.Msgbox("Error", err.Error())
	}
	q.Enqueue(message)
	sender := STDSender{}
	sender.StartOutputting(q)
	time.Sleep(time.Millisecond * 1000)
	pos, err = plat.Calculate(kinematics.D2r(float64(-20)), 0, 0, 0, 0, 0)
	if err != nil {
		log.Println(err)
		g.Msgbox("Error", err.Error())
	}
	message, err = proto.Marshal(pos)
	if err != nil {
		log.Println(err)
		g.Msgbox("Error", err.Error())
	}
	q.Enqueue(message)
	sender.StartOutputting(q)
	time.Sleep(time.Millisecond * 1000)
	g.Msgbox("Tests", "Yaw Test completed")
}
func TestPlatformPitch() {
	q := queue.NewChannelQueue()
	plat := kinematics.NewStewartPlatform(BASE_RADIUS, PLATFORM_RADIUS, HALF_ANGLE_BETWEEN_BASE, HALF_ANGLE_BETWEEN_PLATFORM, SERVO_HORN_LENGTH, ROD_LENGTH, 0)

	pos, err := plat.Calculate(0, 0, kinematics.D2r(float64(24)), 0, 0, 0)
	if err != nil {
		log.Println(err)
		g.Msgbox("Error", err.Error())

	}
	message, err := proto.Marshal(pos)
	if err != nil {
		log.Println(err)
		g.Msgbox("Error", err.Error())

	}
	q.Enqueue(message)
	sender := STDSender{}
	sender.StartOutputting(q)
	time.Sleep(time.Millisecond * 1000)

	pos, err = plat.Calculate(0, 0, kinematics.D2r(float64(-22)), 0, 0, 0)
	if err != nil {
		log.Println(err)
		g.Msgbox("Error", err.Error())
	}
	message, err = proto.Marshal(pos)
	if err != nil {
		log.Println(err)
		g.Msgbox("Error", err.Error())
	}
	q.Enqueue(message)
	sender.StartOutputting(q)
	time.Sleep(time.Millisecond * 1000)
	g.Msgbox("Tests", "Pitch Test completed")
}
func TestPlatformXTrans() {
	q := queue.NewChannelQueue()
	plat := kinematics.NewStewartPlatform(BASE_RADIUS, PLATFORM_RADIUS, HALF_ANGLE_BETWEEN_BASE, HALF_ANGLE_BETWEEN_PLATFORM, SERVO_HORN_LENGTH, ROD_LENGTH, 0)

	pos, err := plat.Calculate(0, 0, 0, float64(6), 0, 0)
	if err != nil {
		log.Println(err)
		g.Msgbox("Error", err.Error())

	}
	message, err := proto.Marshal(pos)
	if err != nil {
		log.Println(err)
		g.Msgbox("Error", err.Error())
	}
	q.Enqueue(message)
	sender := STDSender{}
	sender.StartOutputting(q)
	time.Sleep(time.Millisecond * 1000)
	pos, err = plat.Calculate(0, 0, 0, float64(-5), 0, 0)
	if err != nil {
		log.Println(err)
		g.Msgbox("Error", err.Error())

	}
	message, err = proto.Marshal(pos)
	if err != nil {
		log.Println(err)
		g.Msgbox("Error", err.Error())

	}
	q.Enqueue(message)
	sender.StartOutputting(q)
	time.Sleep(time.Millisecond * 1000)
	g.Msgbox("Tests", "Completed X translation test")
}
func TestPlatformYTrans() {
	q := queue.NewChannelQueue()
	plat := kinematics.NewStewartPlatform(BASE_RADIUS, PLATFORM_RADIUS, HALF_ANGLE_BETWEEN_BASE, HALF_ANGLE_BETWEEN_PLATFORM, SERVO_HORN_LENGTH, ROD_LENGTH, 0)

	pos, err := plat.Calculate(0, 0, 0, 0, float64(5), 0)
	if err != nil {
		log.Println(err)
		g.Msgbox("Error", err.Error())
	}
	message, err := proto.Marshal(pos)
	if err != nil {
		log.Println(err)
		g.Msgbox("Error", err.Error())

	}
	q.Enqueue(message)
	sender := STDSender{}
	sender.StartOutputting(q)
	time.Sleep(time.Millisecond * 1000)
	pos, err = plat.Calculate(0, 0, 0, 0, float64(-5), 0)
	if err != nil {
		log.Println(err)
		g.Msgbox("Error", err.Error())
	}
	message, err = proto.Marshal(pos)
	if err != nil {
		log.Println(err)
		g.Msgbox("Error", err.Error())
	}
	q.Enqueue(message)
	sender.StartOutputting(q)
	time.Sleep(time.Millisecond * 1000)
	g.Msgbox("Tests", "Y translation test completed")
}
func TestPlatformZTrans() {
	q := queue.NewChannelQueue()
	plat := kinematics.NewStewartPlatform(BASE_RADIUS, PLATFORM_RADIUS, HALF_ANGLE_BETWEEN_BASE, HALF_ANGLE_BETWEEN_PLATFORM, SERVO_HORN_LENGTH, ROD_LENGTH, 0)

	pos, err := plat.Calculate(0, 0, 0, 0, 0, float64(3))
	if err != nil {
		log.Println(err)
		g.Msgbox("Error", err.Error())
	}
	message, err := proto.Marshal(pos)
	if err != nil {
		log.Println(err)
		g.Msgbox("Error", err.Error())
	}
	q.Enqueue(message)
	sender := STDSender{}
	sender.StartOutputting(q)
	time.Sleep(time.Millisecond * 1000)
	pos, err = plat.Calculate(0, 0, 0, 0, 0, float64(-5))
	if err != nil {
		log.Println(err)
		g.Msgbox("Error", err.Error())
	}
	message, err = proto.Marshal(pos)
	if err != nil {
		log.Println(err)
		g.Msgbox("Error", err.Error())
	}
	q.Enqueue(message)
	sender.StartOutputting(q)
	time.Sleep(time.Millisecond * 1000)
	g.Msgbox("Tests", "Z translation completed")
}
