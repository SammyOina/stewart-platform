package kinematics

import (
	"log"
	"time"

	g "github.com/AllenDang/giu"
	"github.com/golang/protobuf/proto"
	"github.com/sammyoina/stewart-platform-ui/pipeline"
	"github.com/sammyoina/stewart-platform-ui/queue"
)

func TestPlatformRoll() {
	q := queue.NewChannelQueue()
	plat := NewStewartPlatform(BASE_RADIUS, PLATFORM_RADIUS, HALF_ANGLE_BETWEEN_BASE, HALF_ANGLE_BETWEEN_PLATFORM, SERVO_HORN_LENGTH, ROD_LENGTH, 0)

	for i := -34; i <= 34; i++ {
		pos, err := plat.Calculate(0, d2r(float64(i)), 0, 0, 0, 0)
		if err != nil {
			log.Println(err)
			g.Msgbox("Error", err.Error())
			continue
		}
		message, err := proto.Marshal(pos)
		if err != nil {
			log.Println(err)
			g.Msgbox("Error", err.Error())
			continue
		}
		q.Enqueue(message)
		sender := pipeline.STDSender{}
		sender.StartOutputting(q)
		time.Sleep(time.Millisecond * 200)
	}
}
func TestPlatformYaw() {
	q := queue.NewChannelQueue()
	plat := NewStewartPlatform(BASE_RADIUS, PLATFORM_RADIUS, HALF_ANGLE_BETWEEN_BASE, HALF_ANGLE_BETWEEN_PLATFORM, SERVO_HORN_LENGTH, ROD_LENGTH, 0)

	for i := -20; i <= 20; i++ {
		pos, err := plat.Calculate(d2r(float64(i)), 0, 0, 0, 0, 0)
		if err != nil {
			log.Println(err)
			g.Msgbox("Error", err.Error())
			continue
		}
		message, err := proto.Marshal(pos)
		if err != nil {
			log.Println(err)
			g.Msgbox("Error", err.Error())
			continue
		}
		q.Enqueue(message)
		sender := pipeline.STDSender{}
		sender.StartOutputting(q)
		time.Sleep(time.Millisecond * 200)
	}
}
func TestPlatformPitch() {
	q := queue.NewChannelQueue()
	plat := NewStewartPlatform(BASE_RADIUS, PLATFORM_RADIUS, HALF_ANGLE_BETWEEN_BASE, HALF_ANGLE_BETWEEN_PLATFORM, SERVO_HORN_LENGTH, ROD_LENGTH, 0)

	for i := -22; i <= 24; i++ {
		pos, err := plat.Calculate(0, 0, d2r(float64(i)), 0, 0, 0)
		if err != nil {
			log.Println(err)
			g.Msgbox("Error", err.Error())
			continue
		}
		message, err := proto.Marshal(pos)
		if err != nil {
			log.Println(err)
			g.Msgbox("Error", err.Error())
			continue
		}
		q.Enqueue(message)
		sender := pipeline.STDSender{}
		sender.StartOutputting(q)
		time.Sleep(time.Millisecond * 200)
	}
}
func TestPlatformXTrans() {
	q := queue.NewChannelQueue()
	plat := NewStewartPlatform(BASE_RADIUS, PLATFORM_RADIUS, HALF_ANGLE_BETWEEN_BASE, HALF_ANGLE_BETWEEN_PLATFORM, SERVO_HORN_LENGTH, ROD_LENGTH, 0)

	for i := -5; i <= 6; i++ {
		pos, err := plat.Calculate(0, 0, 0, float64(i), 0, 0)
		if err != nil {
			log.Println(err)
			g.Msgbox("Error", err.Error())
			continue
		}
		message, err := proto.Marshal(pos)
		if err != nil {
			log.Println(err)
			g.Msgbox("Error", err.Error())
			continue
		}
		q.Enqueue(message)
		sender := pipeline.STDSender{}
		sender.StartOutputting(q)
		time.Sleep(time.Millisecond * 200)
	}
}
func TestPlatformYTrans() {
	q := queue.NewChannelQueue()
	plat := NewStewartPlatform(BASE_RADIUS, PLATFORM_RADIUS, HALF_ANGLE_BETWEEN_BASE, HALF_ANGLE_BETWEEN_PLATFORM, SERVO_HORN_LENGTH, ROD_LENGTH, 0)

	for i := -5; i <= 5; i++ {
		pos, err := plat.Calculate(0, 0, 0, 0, float64(i), 0)
		if err != nil {
			log.Println(err)
			g.Msgbox("Error", err.Error())
			continue
		}
		message, err := proto.Marshal(pos)
		if err != nil {
			log.Println(err)
			g.Msgbox("Error", err.Error())
			continue
		}
		q.Enqueue(message)
		sender := pipeline.STDSender{}
		sender.StartOutputting(q)
		time.Sleep(time.Millisecond * 200)
	}
}
func TestPlatformZTrans() {
	q := queue.NewChannelQueue()
	plat := NewStewartPlatform(BASE_RADIUS, PLATFORM_RADIUS, HALF_ANGLE_BETWEEN_BASE, HALF_ANGLE_BETWEEN_PLATFORM, SERVO_HORN_LENGTH, ROD_LENGTH, 0)

	for i := -5; i <= 3; i++ {
		pos, err := plat.Calculate(0, 0, 0, 0, 0, float64(i))
		if err != nil {
			log.Println(err)
			g.Msgbox("Error", err.Error())
			continue
		}
		message, err := proto.Marshal(pos)
		if err != nil {
			log.Println(err)
			g.Msgbox("Error", err.Error())
			continue
		}
		q.Enqueue(message)
		sender := pipeline.STDSender{}
		sender.StartOutputting(q)
		time.Sleep(time.Millisecond * 200)
	}
}
