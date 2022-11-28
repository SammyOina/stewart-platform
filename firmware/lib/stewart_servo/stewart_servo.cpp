#include <stewart_servo.h>
stewart_servo::stewart_servo(int _servo_pins[6], float _delay, int _step_size)
{
    this->delay_duration = _delay;
    this->step_size = _step_size;
    for(int i = 0; i < 6; i++){
        this->servo_pins[i] = _servo_pins[i];
        this->inverted_servo[i] = false;
    }
}

void stewart_servo::init(){
    for(int i = 0; i < 6; i++){
        this->servos[i].attach(this->servo_pins[i]);
        //this->currentAngles[i] = this->servos[i].read();
        this->currentAngles[i] = this->servos[i].read();
        this->targetAngles[i] = 90;
    }
}

void stewart_servo::set_target_angles(float targets[6]){
    for(int i = 0; i < 6; i++){
        this->targetAngles[i] = targets[i];
    }
}
bool stewart_servo::drive(){
    if(this->currentAngles[0] != this->targetAngles[0] || this->currentAngles[1] != this->targetAngles[1] ||
    this->currentAngles[2] != this->targetAngles[2] || this->currentAngles[3] != this->targetAngles[3] ||
    this->currentAngles[4] != this->targetAngles[4] || this->currentAngles[5] != this->targetAngles[5]) {
        for(int i = 0; i < 6; i++){
            float diff = this->targetAngles[i] - this->currentAngles[i];
            if (diff == 0){
                return true;
            }else{
                if (abs(diff) < this->step_size){
                    this->currentAngles[i] = this->targetAngles[i];
                }else{
                    if (diff > 0){
                        this->currentAngles[i] = this->currentAngles[i] + this->step_size;
                    }else{
                        this->currentAngles[i] = this->currentAngles[i] - this->step_size;
                    }
                }
                if (this->inverted_servo[i] == false){
                    this->servos[i].write(this->currentAngles[i]);
                }else{
                    this->servos[i].write(180-this->currentAngles[i]);
                }
            }
        }
        if(this->delay_duration != 0){
            delay(delay_duration);
        }
        return false;
    }else{
        return true;
    }

}

void stewart_servo::invert_servo(int number){
    this->inverted_servo[number] = true;
}