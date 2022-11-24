#ifndef STEWART_SERVO_H
#define STEWART_SERVO_H

#include <ESP32Servo.h>

class stewart_servo
{
private:
    float currentAngles[6];
    float targetAngles[6];
    Servo servos[6];
    int servo_pins[6];
    bool inverted_servo[6];

public:
    float delay_duration;
    int step_size;
    stewart_servo(int _servo_pins[6], float _delay, int _step_size);
    void init();
    void set_target_angles(float targets[6]);
    bool drive();
    void invert_servo(int number);
};


#endif