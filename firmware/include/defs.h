#ifndef DEFS_H
#define DEFS_H
#define STRAIN_CLK 5
#define STRAIN_1 26
#define STRAIN_2 27
#define STRAIN_3 13
#define STRAIN_4 25
#define STRAIN_5 12
#define STRAIN_6 14
#define TARE_TIMEOUT 4
bool mpuActive;
bool intakePitotActive;
bool diffuserPitotActive;
bool positioningActive;
byte DOUTS[6] = {STRAIN_1, STRAIN_2, STRAIN_3, STRAIN_4, STRAIN_5, STRAIN_6};

#define CHANNEL_COUNT sizeof(DOUTS) / sizeof(byte)

long int strain_results[CHANNEL_COUNT];
const char *addr = "192.168.137.1";
const uint16_t port = 8080;
unsigned long timer = 0;
uint8_t buffer[128];
float angRec[6];
bool status;
byte AddressIntake = 0x02;
byte AddressDiffuser = 0x01;
#endif