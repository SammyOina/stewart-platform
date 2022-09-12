package kinematics

import (
	"github.com/sammyoina/stewart-platform-ui/models"
	"github.com/sammyoina/stewart-platform-ui/pipeline"
)

func SetOrientation() {
	position := models.ServoPositionEvent{
		Servo1: 45,
		Servo2: 45,
		Servo3: 45,
		Servo4: 45,
		Servo5: 45,
		Servo6: 45,
	}
	pipeline.ServoPositionChannel <- position
	pipeline.Wg.Add(1)
}
