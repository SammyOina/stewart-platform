int pressure;
int pressSensor = A0;

void setup() {
  // put your setup code here, to run once:
pinMode(pressSensor, INPUT);
Serial.begin(9600);
}

void loop() {
  // put your main code here, to run repeatedly:
pressure = analogRead(pressSensor);
pressure = map(pressure, 0, 1023, 0, 100000); //100000Pa is an assumed value
Serial.println(pressure);
delay(300);
}
