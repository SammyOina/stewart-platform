#include <Arduino.h>
#include "WiFi.h"
#include <WebSocketsClient.h>
#include <events.pb.h>
#include <pb.h>
#include <pb_common.h>
#include <pb_decode.h>
#include <pb_encode.h>
#include "servos.h"

const char *ssid = "sammy2";
const char *password = "12345678";
const char *addr = "192.168.137.1";
const uint16_t port = 8080;

WebSocketsClient webSocket;

uint8_t buffer[128];
bool status;

//QueueHandle_t servoPositionQueue;

//void TaskServoWriter(void * pvParameters);

void printRes(uint8_t *payload, size_t len) {
	ServoPositionEvent message = ServoPositionEvent_init_zero;
	pb_istream_t stream = pb_istream_from_buffer(payload, len);
    status = pb_decode(&stream, ServoPositionEvent_fields, &message);
	/*for(int i = 0; i<len; i++){
    	Serial.printf("%02X",payload[i]);
  	}*/
	if (!status)
	{
		printf("Decoding failed: %s\n", PB_GET_ERROR(&stream));
		return;
    }else{
		for (int i=0; i < 6; i++){
				float angle;
				switch (i){
					case 0:
						angle = message.servo1;
						break;
					case 1:
						angle = message.servo2;
						break;
					case 2:
						angle = message.servo3;
						break;
					case 3:
						angle = message.servo4;
						break;
					case 4:
						angle = message.servo5;
						break;
					case 5:
						angle = message.servo6;
						break;
				}
				Serial.println(angle);
 				if (i != 0 && i != 1 && i != 5) {
 					WriteServoPosition(i, angle, false);
 				}else{
 					WriteServoPosition(i, angle, true);
 				}
 			}
	}
}

void webSocketEvent(WStype_t type, uint8_t * payload, size_t length) {

	switch(type) {
		case WStype_DISCONNECTED:
			Serial.printf("[WSc] Disconnected!\n");
			break;
		case WStype_CONNECTED:
			Serial.printf("[WSc] Connected to url: %s\n", payload);

			// send message to server when Connected
			//webSocket.sendTXT("Connected");
			//webSocket.sendTXT("hey");
			break;
		case WStype_TEXT:
			Serial.printf("[WSc] get text: %s\n", payload);
			//xQueueSend(servoPositionQueue, &payload, portMAX_DELAY);

			//ServoPositionEvent message = ServoPositionEvent_init_zero;
			//uint8_t buf[128];
        	//pb_istream_t stream = pb_istream_from_buffer(payload, sizeof(payload));
			
        	//status = pb_decode(&stream, ServoPositionEvent_fields, &message);
			//Serial.println(message.servo1);
			// send message to server
			// webSocket.sendTXT("message here");
			
			break;
		case WStype_BIN:
			Serial.printf("[WSc] get binary length: %u\n", length);

			// send data to server
			// webSocket.sendBIN(payload, length);
			printRes(payload, length);
			break;
		case WStype_ERROR:			
		case WStype_FRAGMENT_TEXT_START:
		case WStype_FRAGMENT_BIN_START:
		case WStype_FRAGMENT:
		case WStype_FRAGMENT_FIN:
			break;
	}

}

void setup()
{
    Serial.begin(115200);
    WiFi.mode(WIFI_STA); // Make this the client (the server is WIFI_AP)

    delay(100);

	//servo init
	AttachServos();

    WiFi.begin(ssid, password);

    while (WiFi.status() != WL_CONNECTED)
    {
        Serial.println("WIFI connection failed, reconnecting...");
        delay(500);
    }
    

	//webSocket.setExtraHeaders(0);
    	// server address, port and URL
	webSocket.begin("192.168.137.1", port, "/imu", "");

	// event handler
	webSocket.onEvent(webSocketEvent);

	// try ever 5000 again if connection has failed
	webSocket.setReconnectInterval(5000);

	//servoPositionQueue = xQueueCreate(2, sizeof(buffer));

	//if (servoPositionQueue == NULL) {
	//	Serial.println("Failed to create queue");
	//}

	//xTaskCreate(TaskServoWriter, "Write_servo_task", 128, NULL, 1, NULL);
}

void loop()
{
	SensorEvent test = SensorEvent_init_zero;
	test.which_event = SensorEvent_iMUEvent_tag;
	test.event.iMUEvent.pitch = 5.1;
	test.event.iMUEvent.roll = 2.2;
	test.event.iMUEvent.yaw = 3.3;
    pb_ostream_t stream = pb_ostream_from_buffer(buffer, sizeof(buffer));
    if (!pb_encode(&stream, SensorEvent_fields, &test))
    {
        Serial.println("failed to encode temp proto");
        Serial.println(PB_GET_ERROR(&stream));
        return;
    }
	
	webSocket.loop();
	if (webSocket.isConnected()){
		//Serial.println("sending message...");
		webSocket.sendBIN(buffer, stream.bytes_written);
		delay(1000);
	}
}

/*void TaskServoWriter(void * pvParameters){
	uint8_t buf[128];
	bool status;
	while (true) {
		if (xQueueReceive(servoPositionQueue, &buffer, portMAX_DELAY) == pdPASS) {
			ServoPositionEvent message = ServoPositionEvent_init_zero;
        
        	pb_istream_t stream = pb_istream_from_buffer(buf, sizeof(buf));
        
        	status = pb_decode(&stream, ServoPositionEvent_fields, &message);
			Serial.println(message.servo1);
        
	        if (!status)
    	    {
        	    printf("Decoding failed: %s\n", PB_GET_ERROR(&stream));
            	continue;
        	}
			for (int i=0; i < 6; i++){
				float angle;
				switch (i){
					case 0:
						angle = message.servo1;
					case 1:
						angle = message.servo2;
					case 2:
						angle = message.servo3;
					case 3:
						angle = message.servo4;
					case 4:
						angle = message.servo5;
					case 5:
						angle = message.servo6;
				}
 				if (i != 0 && i != 4 && i != 5) {
 					WriteServoPosition(i, angle, false);
 				}else{
 					WriteServoPosition(i, angle, true);
 				}
 			}
		}
	}
}*/