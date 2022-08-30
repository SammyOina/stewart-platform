#include <Arduino.h>
#include "WiFi.h"
#include <ArduinoJson.h>
#include <test.pb.h>
#include <pb.h>
#include <pb_common.h>
#include <pb_decode.h>
#include <pb_encode.h>

const char *ssid = "sammy2";
const char *password = "12345678";
const char *addr = "192.168.137.1";
const uint16_t port = 10101;

WiFiClient client;

uint8_t buffer[128];

void setup()
{
    Serial.begin(115200);
    WiFi.mode(WIFI_STA); // Make this the client (the server is WIFI_AP)

    delay(100);

    WiFi.begin(ssid, password);

    while (WiFi.status() != WL_CONNECTED)
    {
        Serial.println("WIFI connection failed, reconnecting...");
        delay(500);
    }

    DynamicJsonDocument doc(1024);

    doc["leg_strain"][0] = 1.302038;
    doc["leg_strain"][1] = 2.302038;
    doc["leg_strain"][2] = 3.302038;
    doc["leg_strain"][3] = 4.302038;
    doc["leg_strain"][4] = 5.302038;
    doc["leg_strain"][5] = 6.302038;

    serializeJson(doc, buffer);

    TestMessage test_m = TestMessage_init_zero;
    test_m.test_number = 12;
}

void loop()
{
    TestMessage test_m = TestMessage_init_zero;
    test_m.test_number = 12;
    if (!client.connect(addr, port))
    {
        Serial.println("connection failed");
        Serial.println("wait 5 sec to reconnect...");
        delay(5000);
        return;
    }
    pb_ostream_t stream = pb_ostream_from_buffer(buffer, sizeof(buffer));
    if (!pb_encode(&stream, TestMessage_fields, &test_m))
    {
        Serial.println("failed to encode temp proto");
        Serial.println(PB_GET_ERROR(&stream));
        return;
    }

    Serial.print("sending message...");
    Serial.println(test_m.test_number);
    client.write(buffer, stream.bytes_written);
}