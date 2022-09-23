package kinematics

import (
	"testing"
)

func TestRollRange(t *testing.T) {
	plat := NewStewartPlatform(BASE_RADIUS, PLATFORM_RADIUS, HALF_ANGLE_BETWEEN_BASE, HALF_ANGLE_BETWEEN_PLATFORM, SERVO_HORN_LENGTH, ROD_LENGTH, 0)

	for i := -34; i <= 34; i++ {
		_, err := plat.Calculate(0, d2r(float64(i)), 0, 0, 0, 0)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestPitchRange(t *testing.T) {
	plat := NewStewartPlatform(BASE_RADIUS, PLATFORM_RADIUS, HALF_ANGLE_BETWEEN_BASE, HALF_ANGLE_BETWEEN_PLATFORM, SERVO_HORN_LENGTH, ROD_LENGTH, 0)

	for i := -22; i <= 24; i++ {
		_, err := plat.Calculate(0, 0, d2r(float64(i)), 0, 0, 0)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestYawRange(t *testing.T) {
	plat := NewStewartPlatform(BASE_RADIUS, PLATFORM_RADIUS, HALF_ANGLE_BETWEEN_BASE, HALF_ANGLE_BETWEEN_PLATFORM, SERVO_HORN_LENGTH, ROD_LENGTH, 0)

	for i := -20; i <= 20; i++ {
		_, err := plat.Calculate(d2r(float64(i)), 0, 0, 0, 0, 0)
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
