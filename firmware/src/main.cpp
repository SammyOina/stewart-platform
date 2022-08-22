#include <Arduino.h>
#include "WiFi.h"
#include "AsyncUDP.h"

const char * ssid = "stewart";
const char * password = "12345678";

AsyncUDP udp;
IPAddress local_ip(192, 168, 1, 1);
IPAddress gateway(192, 168, 1, 1);
IPAddress subnet(255, 255, 255, 0);

int port = 1234;

void setup() {
    Serial.begin(115200);
    WiFi.mode(WIFI_AP); // Make this the client (the server is WIFI_AP)
    WiFi.softAP(ssid, password);
    WiFi.softAPConfig(local_ip, gateway, subnet);

    delay(100);

    Serial.println(WiFi.softAPIP());

    if(udp.listen(port)) {
        udp.onPacket([](AsyncUDPPacket packet) {
          Serial.print("Received data: ");
            Serial.write(packet.data(), packet.length());
            Serial.println();
        });
    }
}

void loop(){
    delay(1000);
    udp.broadcastTo("Nothing but dreams", port);
}