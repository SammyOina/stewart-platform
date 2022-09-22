int pressure;
int Rx = 0;
int Tx = 1;

void setup() {
  // put your setup code here, to run once:
pinMode(Rx, INPUT);
pinMode(Tx, OUTPUT);
Serial.begin(9600);
}

void loop() {
  // put your main code here, to run repeatedly:
pressure = Serial.available();
Serial.println(pressure);
delay(300);
}
