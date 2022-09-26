#include <ESP32Servo.h>

#define SERVOS 6

Servo myservo[SERVOS];

int servo_pins[SERVOS] = {33,32,19,4,23,18};

void AttachServos(void) {
    for(int i = 0; i < SERVOS; i++) {
        // Attach the servo to the servo object
        myservo[i].attach(servo_pins[i]);
        delay(500);
        myservo[i].write(90);
    }
}

void WriteServoPosition( int servo, float angle, bool reverse) {
    if (reverse == false){
        myservo[servo].write(angle);
    }else{
        myservo[servo].write(180-angle);
    }
}