#include <Arduino.h>
//#include "WiFi.h"
#include <WebSocketsClient.h>
#include <events.pb.h>
#include <pb.h>
#include <pb_common.h>
#include <pb_decode.h>
#include <pb_encode.h>
//#include "servos.h"
#include <SensorModbusMaster.h>
#include "Wire.h"
#include <MPU6050_light.h>
#include <HX711-multi.h>
#include <stewart_servo.h>
#include <WiFiManager.h>

#define STRAIN_CLK 5
#define STRAIN_1 26
#define STRAIN_2 27
#define STRAIN_3 13
#define STRAIN_4 25
#define STRAIN_5 12
#define STRAIN_6 14

WiFiManager wm;

#define TARE_TIMEOUT 4

byte DOUTS[6] = {STRAIN_1, STRAIN_2, STRAIN_3, STRAIN_4, STRAIN_5, STRAIN_6};

#define CHANNEL_COUNT sizeof(DOUTS)/sizeof(byte)

long int strain_results[CHANNEL_COUNT];

HX711MULTI scales(CHANNEL_COUNT, DOUTS, STRAIN_CLK);

const char *ssid = "sammy2";
const char *password = "12345678";
const char *addr = "192.168.137.1";
const uint16_t port = 8080;

MPU6050 mpu(Wire);
unsigned long timer = 0;

WebSocketsClient webSocket;

uint8_t buffer[128];
float angRec[6];
bool status;

HardwareSerial modbusSerial = Serial2;

modbusMaster intakePitot;
modbusMaster diffuserPitot;

byte AddressIntake = 0x02;
byte AddressDiffuser = 0x01;

QueueHandle_t servoPositionQueue;

void TaskServoWriter(void * pvParameters);

void tare() {
  bool tareSuccessful = false;

  unsigned long tareStartTime = millis();
  while (!tareSuccessful && millis()<(tareStartTime+TARE_TIMEOUT*1000)) {
    tareSuccessful = scales.tare(20,10000);  //reject 'tare' if still ringing
  }
}

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
				angRec[i] = angle;
				/*Serial.println(angle);
 				if (i != 0 && i != 2 && i != 4) {
 					WriteServoPosition(i, angle, false);
 				}else{
 					WriteServoPosition(i, angle, true);
 				}*/
 			}
		xQueueSend(servoPositionQueue, &angRec, portMAX_DELAY);

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
	//tare();
    WiFi.mode(WIFI_STA); // Make this the client (the server is WIFI_AP)

    delay(100);

	wm.setConfigPortalBlocking(false);
	if(wm.autoConnect("AutoConnectAP")){
        Serial.println("connected...yeey :)");
    }
    else {
        Serial.println("Configportal running");
    }

	//servo init
	//AttachServos();

    //WiFi.begin(ssid, password);

    /*while (WiFi.status() != WL_CONNECTED)
    {*/
        //Serial.println("WIFI connection failed, reconnecting...");
        //delay(500);
    //}
    

	//webSocket.setExtraHeaders(0);
    	// server address, port and URL
	webSocket.begin("192.168.137.1", port, "/imu", "");

	// event handler
	webSocket.onEvent(webSocketEvent);

	// try ever 5000 again if connection has failed
	webSocket.setReconnectInterval(5000);

	servoPositionQueue = xQueueCreate(2, sizeof(angRec));

	if (servoPositionQueue == NULL) {
		Serial.println("Failed to create queue");
	}

	xTaskCreatePinnedToCore(TaskServoWriter, "Write_servo_task", 2048, NULL, 1, NULL, 1);
	modbusSerial.begin(9600);
	intakePitot.begin(AddressIntake, modbusSerial);
	diffuserPitot.begin(AddressDiffuser, modbusSerial);

	Wire.begin();
	byte status = mpu.begin();
  	Serial.print(F("MPU6050 status: "));
  	Serial.println(status);
  	while(status!=0){ } // stop everything if could not connect to MPU6050
  
  	Serial.println(F("Calculating offsets, do not move MPU6050"));
  	delay(1000);
  	// mpu.upsideDownMounting = true; // uncomment this line if the MPU6050 is mounted upside-down
  	mpu.calcOffsets(); // gyro and accelero
  	Serial.println("Done!\n");
	tare();
}

void loop()
{
	mpu.update();
	if ((millis()-timer)>10){
		SensorEvent orientation = SensorEvent_init_zero;
		orientation.which_event = SensorEvent_iMUEvent_tag;
		orientation.event.iMUEvent.pitch = mpu.getAngleX();
		orientation.event.iMUEvent.roll = mpu.getAngleZ();
		orientation.event.iMUEvent.yaw = mpu.getAngleY();
    	pb_ostream_t stream = pb_ostream_from_buffer(buffer, sizeof(buffer));
    	if (!pb_encode(&stream, SensorEvent_fields, &orientation))
    	{
        	Serial.println("failed to encode temp proto");
        	Serial.println(PB_GET_ERROR(&stream));
        	return;
    	}
	
		if (webSocket.isConnected()){
			//Serial.println("sending message...");
			webSocket.sendBIN(buffer, stream.bytes_written);
			//delay(1000);
		}
		timer = millis();
	}
	delay(100);
	SensorEvent pitotReading = SensorEvent_init_zero;
		pitotReading.which_event = SensorEvent_pitotEvent_tag;
		pitotReading.event.pitotEvent.intakePitot = 5;
		pitotReading.event.pitotEvent.diffuserPitot = 3;
		pitotReading.event.pitotEvent.testSectionPitot = 2;

		pb_ostream_t stream2 = pb_ostream_from_buffer(buffer, sizeof(buffer));
    	if (!pb_encode(&stream2, SensorEvent_fields, &pitotReading))
    	{
        	Serial.println("failed to encode temp proto");
        	Serial.println(PB_GET_ERROR(&stream2));
        	return;
    	}
		if (webSocket.isConnected()){
			//Serial.println("sending message...");
			webSocket.sendBIN(buffer, stream2.bytes_written);
			//delay(1000);
		}
		delay(100);
		/*SensorEvent strains = SensorEvent_init_zero;
		strains.which_event = SensorEvent_strainEvent_tag;
		strains.event.strainEvent.strain1 = 0.5;
		strains.event.strainEvent.strain2 = 1.5;
		strains.event.strainEvent.strain3 = 2.5;
		strains.event.strainEvent.strain4 = 3.5;
		strains.event.strainEvent.strain5 = 6.5;
		strains.event.strainEvent.strain6 = 7.5;
		pb_ostream_t stream1 = pb_ostream_from_buffer(buffer, sizeof(buffer));
    	if (!pb_encode(&stream1, SensorEvent_fields, &strains))
    	{
    		Serial.println("failed to encode temp proto");
    		Serial.println(PB_GET_ERROR(&stream1));
    		return;
    	}
		if (webSocket.isConnected()){
			//Serial.println("sending message...");
			webSocket.sendBIN(buffer, stream1.bytes_written);
			//delay(1000);
		}
		delay(100);*/
	
	/*bool gotReadingIntake = intakePitot.getRegisters(0x03, 0x00, 3);
	bool gotReadingDiffuser = diffuserPitot.getRegisters(0x03, 0x00,3);
	if (gotReadingIntake && gotReadingDiffuser) {
		SensorEvent pitotReading = SensorEvent_init_zero;
		pitotReading.which_event = SensorEvent_pitotEvent_tag;
		pitotReading.event.pitotEvent.intakePitot = intakePitot.int16FromFrame(bigEndian, 3);
		pitotReading.event.pitotEvent.diffuserPitot = diffuserPitot.int16FromFrame(bigEndian, 3);

		pb_ostream_t stream = pb_ostream_from_buffer(buffer, sizeof(buffer));
    	if (!pb_encode(&stream, SensorEvent_fields, &pitotReading))
    	{
        	Serial.println("failed to encode temp proto");
        	Serial.println(PB_GET_ERROR(&stream));
        	return;
    	}
		if (webSocket.isConnected()){
			//Serial.println("sending message...");
			webSocket.sendBIN(buffer, stream.bytes_written);
			//delay(1000);
		}
	}*/
	if (scales.is_ready()){
		scales.read(strain_results);
		for (int i = 0; i <6; i++){
			Serial.print("strain");
			Serial.print(i);
			Serial.print(" ");
			Serial.println(strain_results[i]);
			delay(100);
		}
		SensorEvent strains = SensorEvent_init_zero;
		strains.which_event = SensorEvent_strainEvent_tag;
		strains.event.strainEvent.strain1 = strain_results[0];
		strains.event.strainEvent.strain2 = strain_results[1];
		strains.event.strainEvent.strain3 = strain_results[2];
		strains.event.strainEvent.strain4 = strain_results[3];
		strains.event.strainEvent.strain5 = strain_results[4];
		strains.event.strainEvent.strain6 = strain_results[5];
		pb_ostream_t stream = pb_ostream_from_buffer(buffer, sizeof(buffer));
    	if (!pb_encode(&stream, SensorEvent_fields, &strains))
    	{
    		Serial.println("failed to encode temp proto");
    		Serial.println(PB_GET_ERROR(&stream));
    		return;
    	}
		if (webSocket.isConnected()){
			//Serial.println("sending message...");
			webSocket.sendBIN(buffer, stream.bytes_written);
			//delay(1000);
		}
		}
	
	webSocket.loop();
	wm.process();
}

void TaskServoWriter(void * pvParameters){
	float angTargets[6];
	int servo_pins[6] = {33,32,19,4,23,18};
	stewart_servo stewartServo(servo_pins, 20,1);
	stewartServo.invert_servo(0);
	stewartServo.invert_servo(2);
	stewartServo.invert_servo(4);
	stewartServo.init();

	while (true) {
		if (xQueueReceive(servoPositionQueue, &angTargets, 1000) == pdPASS) {
			stewartServo.set_target_angles(angTargets);
			Serial.println("new position");
			/*for (int i = 0; i < 6; i++){
				if (i != 0 && i != 2 && i != 4) {
 					WriteServoPosition(i, angTargets[i], false);
 				}else{
 					WriteServoPosition(i, angTargets[i], true);
 				}
			}*/
		}
		if (stewartServo.drive() != true){
			Serial.println("moving to position");
			tare();
		}else{
			Serial.println("in position");
		}
	}
}