#include <Arduino.h>
#include <WebSocketsClient.h>
#include <events.pb.h>
#include <pb.h>
#include <pb_common.h>
#include <pb_decode.h>
#include <pb_encode.h>
#include <SensorModbusMaster.h>
#include "Wire.h"
#include <MPU6050_light.h>
#include <HX711-multi.h>
#include <stewart_servo.h>
#include <WiFiManager.h>
#include <defs.h>

WiFiManager wm;

HX711MULTI scales(CHANNEL_COUNT, DOUTS, STRAIN_CLK);

MPU6050 mpu(Wire);

WebSocketsClient webSocket;

HardwareSerial modbusSerial = Serial2;

modbusMaster intakePitot;
modbusMaster diffuserPitot;

QueueHandle_t servoPositionQueue;

bool testing_mode;

void TaskServoWriter(void *pvParameters);
void configModeCallback(WiFiManager *myWiFiManager);
void tare();
void getServoPosition(uint8_t *payload, size_t len);
void webSocketEvent(WStype_t type, uint8_t *payload, size_t length);

void setup()
{
	testing_mode = false;
	Serial.begin(115200);
	positioningActive = false;

	Serial.println("done here");
	WiFi.mode(WIFI_STA); // Make this the client (the server is WIFI_AP)

	delay(100);

	if (!wm.autoConnect())
	{
		Serial.println("failed to connect and hit timeout");
		// reset and try again, or maybe put it to deep sleep
		ESP.restart();
		delay(1000);
	}
	Serial.println("connected...yeey :)");

	wm.setAPCallback(configModeCallback);
	//  server address, port and URL
	webSocket.begin("192.168.137.1", port, "/imu", "");

	// event handler
	webSocket.onEvent(webSocketEvent);

	// try ever 5000 again if connection has failed
	webSocket.setReconnectInterval(5000);

	servoPositionQueue = xQueueCreate(2, sizeof(angRec));

	if (servoPositionQueue == NULL)
	{
		Serial.println("Failed to create queue");
	}

	modbusSerial.begin(9600);
	intakePitotActive = intakePitot.begin(AddressIntake, modbusSerial);
	diffuserPitotActive = diffuserPitot.begin(AddressDiffuser, modbusSerial);

	Wire.begin();
	byte status = mpu.begin();
	Serial.print(F("MPU6050 status: "));
	Serial.println(status);
	if (status == 0)
	{
		mpuActive = true;
	}
	/*while (status != 0)
	{
	} // stop everything if could not connect to MPU6050*/

	Serial.println(F("Calculating offsets, do not move MPU6050"));
	delay(1000);
	// mpu.upsideDownMounting = true; // uncomment this line if the MPU6050 is mounted upside-down
	mpu.calcOffsets(); // gyro and accelero
	Serial.println("Done!\n");
	scales.setDebugEnable();
	tare();
	xTaskCreatePinnedToCore(TaskServoWriter, "Write_servo_task", 2048, NULL, 1, NULL, 1);
}

void loop()
{
	webSocket.loop();
	if (testing_mode)
	{
		SensorEvent orientation = SensorEvent_init_zero;
		orientation.which_event = SensorEvent_iMUEvent_tag;
		orientation.event.iMUEvent.pitch = 0;
		orientation.event.iMUEvent.roll = 0;
		orientation.event.iMUEvent.yaw = 0;
		pb_ostream_t stream = pb_ostream_from_buffer(buffer, sizeof(buffer));
		if (!pb_encode(&stream, SensorEvent_fields, &orientation))
		{
			Serial.println("failed to encode temp proto");
			Serial.println(PB_GET_ERROR(&stream));
			return;
		}

		if (webSocket.isConnected())
		{
			Serial.println("sending message...");
			webSocket.sendBIN(buffer, stream.bytes_written);
			// delay(1000);
		}
		SensorEvent strains = SensorEvent_init_zero;
		strains.which_event = SensorEvent_strainEvent_tag;
		strains.event.strainEvent.strain1 = 10;
		strains.event.strainEvent.strain2 = 10;
		strains.event.strainEvent.strain3 = 10;
		strains.event.strainEvent.strain4 = 10;
		strains.event.strainEvent.strain5 = 10;
		strains.event.strainEvent.strain6 = 10;
		stream = pb_ostream_from_buffer(buffer, sizeof(buffer));
		if (!pb_encode(&stream, SensorEvent_fields, &strains))
		{
			Serial.println("failed to encode temp proto");
			Serial.println(PB_GET_ERROR(&stream));
			return;
		}
		if (webSocket.isConnected())
		{
			Serial.println("sending message...");
			webSocket.sendBIN(buffer, stream.bytes_written);
			// delay(1000);
		}
		delay(100);
	}
	else
	{
		if (!positioningActive)
		{
			if (mpuActive)
			{
				mpu.update();
				if ((millis() - timer) > 10)
				{
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

					if (webSocket.isConnected())
					{
						// Serial.println("sending message...");
						webSocket.sendBIN(buffer, stream.bytes_written);
						// delay(1000);
					}
					timer = millis();
				}
			}
			if (scales.is_ready())
			{
				scales.read(strain_results);
				// Serial.println("reading");
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
				if (webSocket.isConnected())
				{
					// Serial.println("sending message...");
					webSocket.sendBIN(buffer, stream.bytes_written);
					// delay(1000);
				}
			}
		}
	}
}

void TaskServoWriter(void *pvParameters)
{
	float angTargets[6];
	int servo_pins[6] = {33, 32, 19, 4, 23, 18};
	stewart_servo stewartServo(servo_pins, 100, 1);
	stewartServo.invert_servo(0);
	stewartServo.invert_servo(2);
	stewartServo.invert_servo(4);
	stewartServo.init();

	while (true)
	{
		if (xQueueReceive(servoPositionQueue, &angTargets, 1000) == pdPASS)
		{
			stewartServo.set_target_angles(angTargets);
			Serial.println("new position");
			while (!stewartServo.drive())
			{
				Serial.println("moving to position");
			}
			tare();
			positioningActive = false;
		}
		if (testing_mode)
		{
			SensorEvent pitotReading = SensorEvent_init_zero;
			pitotReading.which_event = SensorEvent_pitotEvent_tag;
			pitotReading.event.pitotEvent.intakePitot = 1.2;
			pitotReading.event.pitotEvent.diffuserPitot = 1.1;

			pb_ostream_t stream = pb_ostream_from_buffer(buffer, sizeof(buffer));
			if (!pb_encode(&stream, SensorEvent_fields, &pitotReading))
			{
				Serial.println("failed to encode temp proto");
				Serial.println(PB_GET_ERROR(&stream));
				return;
			}
			if (webSocket.isConnected())
			{
				// Serial.println("sending message...");
				webSocket.sendBIN(buffer, stream.bytes_written);
				// delay(1000);
			}
		}
		else
		{
			if ((intakePitotActive || diffuserPitotActive) && !positioningActive) // will cause disconnection if not connected
			{
				bool gotReadingIntake = intakePitot.getRegisters(0x03, 0x00, 3);
				bool gotReadingDiffuser = diffuserPitot.getRegisters(0x03, 0x00, 3);
				if (gotReadingIntake || gotReadingDiffuser)
				{
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
					if (webSocket.isConnected())
					{
						// Serial.println("sending message...");
						webSocket.sendBIN(buffer, stream.bytes_written);
						// delay(1000);
					}
				}
			}
		}
	}
}
void configModeCallback(WiFiManager *myWiFiManager)
{
	Serial.println("Entered config mode");
	Serial.println(WiFi.softAPIP());
	// if you used auto generated SSID, print it
	Serial.println(myWiFiManager->getConfigPortalSSID());
	// entered config mode, make led toggle faster
}

void tare()
{
	bool tareSuccessful = false;

	unsigned long tareStartTime = millis();
	while (!tareSuccessful && millis() < (tareStartTime + TARE_TIMEOUT * 1000))
	{
		tareSuccessful = scales.tare(20, 10000); // reject 'tare' if still ringing
		Serial.print("tare: ");
		Serial.println(tareSuccessful);
	}
}

void getServoPosition(uint8_t *payload, size_t len)
{
	ServoPositionEvent message = ServoPositionEvent_init_zero;
	pb_istream_t stream = pb_istream_from_buffer(payload, len);
	status = pb_decode(&stream, ServoPositionEvent_fields, &message);
	if (!status)
	{
		printf("Decoding failed: %s\n", PB_GET_ERROR(&stream));
		return;
	}
	else
	{
		for (int i = 0; i < 6; i++)
		{
			float angle;
			switch (i)
			{
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
		}
		xQueueSend(servoPositionQueue, &angRec, portMAX_DELAY);
	}
}

void webSocketEvent(WStype_t type, uint8_t *payload, size_t length)
{

	switch (type)
	{
	case WStype_DISCONNECTED:
		Serial.printf("[WSc] Disconnected!\n");
		break;
	case WStype_CONNECTED:
		Serial.printf("[WSc] Connected to url: %s\n", payload);

		break;
	case WStype_TEXT:
		Serial.printf("[WSc] get text: %s\n", payload);

		break;
	case WStype_BIN:
		Serial.printf("[WSc] get binary length: %u\n", length);
		positioningActive = true;

		getServoPosition(payload, length);
		break;
	case WStype_ERROR:
	case WStype_FRAGMENT_TEXT_START:
	case WStype_FRAGMENT_BIN_START:
	case WStype_FRAGMENT:
	case WStype_FRAGMENT_FIN:
		break;
	}
}