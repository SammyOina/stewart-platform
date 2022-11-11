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

	for i := 0; i <= 34; i++ {
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
	for i := 34; i >= -34; i-- {
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
	g.Msgbox("Tests", "Roll test done")
}
func TestPlatformYaw() {
	q := queue.NewChannelQueue()
	plat := NewStewartPlatform(BASE_RADIUS, PLATFORM_RADIUS, HALF_ANGLE_BETWEEN_BASE, HALF_ANGLE_BETWEEN_PLATFORM, SERVO_HORN_LENGTH, ROD_LENGTH, 0)

	for i := 0; i <= 20; i++ {
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
	for i := 20; i >= -20; i-- {
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
	g.Msgbox("Tests", "Yaw Test completed")
}
func TestPlatformPitch() {
	q := queue.NewChannelQueue()
	plat := NewStewartPlatform(BASE_RADIUS, PLATFORM_RADIUS, HALF_ANGLE_BETWEEN_BASE, HALF_ANGLE_BETWEEN_PLATFORM, SERVO_HORN_LENGTH, ROD_LENGTH, 0)

	for i := 0; i <= 24; i++ {
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
	for i := 24; i >= -22; i-- {
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
	g.Msgbox("Tests", "Pitch Test completed")
}
func TestPlatformXTrans() {
	q := queue.NewChannelQueue()
	plat := NewStewartPlatform(BASE_RADIUS, PLATFORM_RADIUS, HALF_ANGLE_BETWEEN_BASE, HALF_ANGLE_BETWEEN_PLATFORM, SERVO_HORN_LENGTH, ROD_LENGTH, 0)

	for i := 0; i <= 6; i++ {
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
	for i := 6; i >= -5; i-- {
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
	g.Msgbox("Tests", "Completed X translation test")
}
func TestPlatformYTrans() {
	q := queue.NewChannelQueue()
	plat := NewStewartPlatform(BASE_RADIUS, PLATFORM_RADIUS, HALF_ANGLE_BETWEEN_BASE, HALF_ANGLE_BETWEEN_PLATFORM, SERVO_HORN_LENGTH, ROD_LENGTH, 0)

	for i := 0; i <= 5; i++ {
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
	for i := 5; i >= -5; i-- {
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
	g.Msgbox("Tests", "Y translation test completed")
}
func TestPlatformZTrans() {
	q := queue.NewChannelQueue()
	plat := NewStewartPlatform(BASE_RADIUS, PLATFORM_RADIUS, HALF_ANGLE_BETWEEN_BASE, HALF_ANGLE_BETWEEN_PLATFORM, SERVO_HORN_LENGTH, ROD_LENGTH, 0)

	for i := 0; i <= 3; i++ {
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
	for i := 3; i >= -5; i-- {
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
	g.Msgbox("Tests", "Z translation completed")
}
