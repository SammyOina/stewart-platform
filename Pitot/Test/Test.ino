#include <WiFi.h>
const char* ssid = "sammy"
const char* password = "12345678"


int pressure;
int pressSensor = 15;

void setup() {
  // put your setup code here, to run once:
pinMode(pressSensor, INPUT);
Serial.begin(115200);
delay(1000);
WiFi.mode(WIFI_STA);
WiFi.begin(ssid, password);
Serial.println("\n Connecting");

while (WiFi.status() != WL_CONNECTED)
{
  Serial.print(".")
  delay(100);
}
Serial.println("\nConnected to the WiFi network")
{
  Seial.print("Local ESP32 IP: ");
  Serial.println(WiFi.localIP());
}
}

void loop() {
  // put your main code here, to run repeatedly:
pressure = analogRead(pressSensor);
pressure = map(pressure, 0, 1023, -2000, 2000); //1kpa == 0.145psi
Serial.print(pressure);
Serial.println("pa");
delay(300);
}
