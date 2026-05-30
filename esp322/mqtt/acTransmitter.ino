#include <IRremoteESP8266.h>
#include <irsend.h>
#include <ir_Sharp.h>

const uint16_t kIrLed18 = 18;
const uint16_t kIrLed19 = 19;
const uint16_t kIrLed21 = 21;
const int statusLed = 2;

IRSharpAc irsend18(kIrLed18);
IRSharpAc irsend19(kIrLed19);
IRSharpAc irsend21(kIrLed21);

void setupAC() {
    pinMode(statusLed, OUTPUT);
    irsend18.begin();
    irsend19.begin();
    irsend21.begin();
    Serial.println("[AC] IR Transmitters initialized.");
}

void handleAC(String msg) {
    int separator = msg.indexOf(':');
    if (separator == -1) return;
    int acIndex  = msg.substring(0, separator).toInt();
    String hex   = msg.substring(separator + 1);
    if (acIndex == 0) return;
    sendHex(hex, acIndex);
}

void sendHex(String hex, int acIndex) {
    digitalWrite(statusLed, HIGH);
    int byteCount = hex.length() / 2;
    uint8_t state[byteCount];
    for (int i = 0; i < byteCount; i++) {
        state[i] = (uint8_t)strtol(hex.substring(i * 2, i * 2 + 2).c_str(), NULL, 16);
    }
    switch (acIndex) {
        case 1: irsend18.setRaw(state); irsend18.send(); Serial.println("[IR] Sent via Pin 18"); break;
        case 2: irsend19.setRaw(state); irsend19.send(); Serial.println("[IR] Sent via Pin 19"); break;
        case 3: irsend21.setRaw(state); irsend21.send(); Serial.println("[IR] Sent via Pin 21"); break;
        default: Serial.println("[IR] Unknown AC index."); return;
    }
    delay(500);
    digitalWrite(statusLed, LOW);
}