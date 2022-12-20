package kinematics

import (
	"testing"
)

const (
	ROD_LENGTH                  float64 = 14.0
	BASE_RADIUS                 float64 = 15.0
	PLATFORM_RADIUS             float64 = 10.0
	SERVO_HORN_LENGTH           float64 = 4.5
	HALF_ANGLE_BETWEEN_BASE     float64 = 13
	HALF_ANGLE_BETWEEN_PLATFORM float64 = 13
)

func TestRollRange(t *testing.T) {
	plat := NewStewartPlatform(BASE_RADIUS, PLATFORM_RADIUS, HALF_ANGLE_BETWEEN_BASE, HALF_ANGLE_BETWEEN_PLATFORM, SERVO_HORN_LENGTH, ROD_LENGTH, 0)

	for i := -34; i <= 34; i++ {
		_, err := plat.Calculate(0, D2r(float64(i)), 0, 0, 0, 0)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestPitchRange(t *testing.T) {
	plat := NewStewartPlatform(BASE_RADIUS, PLATFORM_RADIUS, HALF_ANGLE_BETWEEN_BASE, HALF_ANGLE_BETWEEN_PLATFORM, SERVO_HORN_LENGTH, ROD_LENGTH, 0)

	for i := -22; i <= 24; i++ {
		_, err := plat.Calculate(0, 0, D2r(float64(i)), 0, 0, 0)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestYawRange(t *testing.T) {
	plat := NewStewartPlatform(BASE_RADIUS, PLATFORM_RADIUS, HALF_ANGLE_BETWEEN_BASE, HALF_ANGLE_BETWEEN_PLATFORM, SERVO_HORN_LENGTH, ROD_LENGTH, 0)

	for i := -20; i <= 20; i++ {
		_, err := plat.Calculate(D2r(float64(i)), 0, 0, 0, 0, 0)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestXTranslationRange(t *testing.T) {
	plat := NewStewartPlatform(BASE_RADIUS, PLATFORM_RADIUS, HALF_ANGLE_BETWEEN_BASE, HALF_ANGLE_BETWEEN_PLATFORM, SERVO_HORN_LENGTH, ROD_LENGTH, 0)

	for i := -5; i <= 6; i++ {
		_, err := plat.Calculate(0, 0, 0, float64(i), 0, 0)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestYTranslationRange(t *testing.T) {
	plat := NewStewartPlatform(BASE_RADIUS, PLATFORM_RADIUS, HALF_ANGLE_BETWEEN_BASE, HALF_ANGLE_BETWEEN_PLATFORM, SERVO_HORN_LENGTH, ROD_LENGTH, 0)

	for i := -5; i <= 5; i++ {
		_, err := plat.Calculate(0, 0, 0, 0, float64(i), 0)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestZTranslationRange(t *testing.T) {
	plat := NewStewartPlatform(BASE_RADIUS, PLATFORM_RADIUS, HALF_ANGLE_BETWEEN_BASE, HALF_ANGLE_BETWEEN_PLATFORM, SERVO_HORN_LENGTH, ROD_LENGTH, 0)

	for i := -5; i <= 3; i++ {
		_, err := plat.Calculate(0, 0, 0, 0, 0, float64(i))
		if err != nil {
			t.Error(err)
		}
	}
}
