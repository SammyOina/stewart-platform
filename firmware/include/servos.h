#include <ESP32Servo.h>

#define SERVOS 6

Servo myservo[SERVOS];

int servo_pins[SERVOS] = {7,8,9,10,11,12};

void AttachServos(void) {
    for(int i = 0; i < SERVOS; i++) {
        // Attach the servo to the servo object
        myservo[i].attach(servo_pins[i]);
        delay(500);
    }
}

void WriteServoPosition( int servo, float angle, bool reverse) {
    if (reverse == false){
        myservo[servo].write(angle);
    }else{
        myservo[servo].write(180-angle);
    }
}