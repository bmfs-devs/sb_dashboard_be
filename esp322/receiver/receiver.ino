#include <Arduino.h>
#include <IRremoteESP8266.h>
#include <IRrecv.h>
#include <IRutils.h>

// Use GPIO 19 as we discussed for receiving
const uint16_t kRecvPin = 19;

// 1024 buffer is plenty, but we enable 'true' for extra precision
IRrecv irrecv(kRecvPin, 1024, 50, true);
decode_results results;

void setup() {
  Serial.begin(115200);
  irrecv.enableIRIn();
  pinMode(2, OUTPUT); 
  Serial.println("--- 104-BIT SHARP AC DECODER READY ---");
}

void loop() {
  if (irrecv.decode(&results)) {
    digitalWrite(2, HIGH); // Blink on reception

    // 1. Print Protocol and Bit Count
    Serial.print("PROTOCOL: ");
    Serial.println(typeToString(results.decode_type));
    Serial.print("BITS: ");
    Serial.println(results.bits);

    // 2. Print Full Hex State (The most important part for your Go backend)
    // If it's a SHARP_AC protocol, we use the state buffer instead of .value
    Serial.print("FULL_HEX: ");
    if (results.decode_type == SHARP_AC || results.bits > 64) {
      // Print the bytes in the state array
      for (uint16_t i = 0; i < results.bits / 8; i++) {
        if (results.state[i] < 0x10) Serial.print("0"); // Leading zero padding
        Serial.print(results.state[i], HEX);
      }
    } else {
      // Fallback for smaller codes
      Serial.print(uint64ToString(results.value, 16));
    }
    Serial.println();

    Serial.println("---");
    
    delay(200);
    digitalWrite(2, LOW);
    irrecv.resume(); // Receive the next signal
  }
}