#include <WiFi.h>
#include <WiFiClientSecure.h> 
#include <PubSubClient.h>
#include <IRremoteESP8266.h>
#include <irsend.h>
#include <ir_Sharp.h>

// --- WiFi Credentials ---
const char* ssid     = "blindspot";
const char* password = "blindspot207";

// --- HiveMQ Credentials ---
const char* mqtt_server = "";
const int mqtt_port     = 8883;
const char* mqtt_user   = "";
const char* mqtt_pass   = "";
const char* mqtt_topic  = "ac_topic";

// --- Pins ---
const uint16_t kIrLed18 = 18;
const uint16_t kIrLed19 = 19;
const uint16_t kIrLed21 = 21;
const int statusLed = 2;

IRSharpAc irsend18(kIrLed18);
IRSharpAc irsend19(kIrLed19);
IRSharpAc irsend21(kIrLed21);
WiFiClientSecure espClient;
PubSubClient client(espClient);

void callback(char* topic, byte* payload, unsigned int length) {
    String msg = "";
    for (int i = 0; i < length; i++) msg += (char)payload[i];
    msg.trim();

    Serial.print("\n[MQTT] Received: ");
    Serial.println(msg);
    int acIndex = 0; 
    String hex = "";
    parseInput(msg, acIndex, hex);
    Serial.println(acIndex);
    Serial.println(hex);

    if (acIndex == 0) {
        return;
    }
    sendHex(hex, acIndex);
}

void parseInput(String input, int &acIndex, String &hex) {
    int separator = input.indexOf(':');

    if (separator != -1) {
       String acIndexStr = input.substring(0, separator);
       acIndex = acIndexStr.toInt();
       hex = input.substring(separator+1);      
    } 
}

void sendHex(String hex, int acIndex) {
    digitalWrite(statusLed, HIGH);
    uint8_t state[13];
    for (int i = 0; i < 13; i++) {
        state[i] = (uint8_t)strtol(hex.substring(i*2, i*2+2).c_str(), NULL, 16);
    }
    switch(acIndex) {
        case 1:
            Serial.println("[IR] Sharp 56kHz Signal Sent. Pin 18");
            irsend18.setRaw(state);
            irsend18.send();
            break;
        case 2:
            Serial.println("[IR] Sharp 56kHz Signal Sent. Pin 19");
            irsend19.setRaw(state);
            irsend19.send();
            break;
        case 3:
            Serial.println("[IR] Sharp 56kHz Signal Sent. Pin 21");
            irsend21.setRaw(state);
            irsend21.send();
            break;
        default:
            Serial.println("[IR] No index found");
            return;
    }
    Serial.println("[IR] Sharp 56kHz Signal Sent.");
    delay(500);
    digitalWrite(statusLed, LOW);
}

void setup_wifi() {
    Serial.println("\n--- WiFi Setup ---");
    WiFi.disconnect(true);
    WiFi.mode(WIFI_STA);
    WiFi.begin(ssid, password);

    int timeout = 0;
    while (WiFi.status() != WL_CONNECTED) {
        delay(1000);
        timeout++;
        Serial.printf("Connecting... Status: %d (Attempt %d)\n", WiFi.status(), timeout);
        
        if (timeout > 20) {
            Serial.println("WiFi Timeout. Check SSID/Password or Signal.");
            return; // Exit and let loop() try again
        }
    }
    Serial.println("WiFi Connected! IP: " + WiFi.localIP().toString());
    espClient.setInsecure(); // Required for HiveMQ TLS
}

void reconnect() {
    while (!client.connected()) {
        if (WiFi.status() != WL_CONNECTED) setup_wifi();
        
        Serial.print("Attempting HiveMQ Connection...");
        String clientId = "ESP32_Client_" + String(random(0xffff), HEX);
        
        if (client.connect(clientId.c_str(), mqtt_user, mqtt_pass)) {
            Serial.println("CONNECTED");
            client.subscribe(mqtt_topic);
        } else {
            Serial.printf("FAILED (rc=%d). Retry in 5s...\n", client.state());
            delay(5000);
        }
    }
}

void setup() {
    Serial.begin(115200);
    pinMode(statusLed, OUTPUT);
    irsend18.begin();
    irsend19.begin();
    irsend21.begin();
    setup_wifi();
    client.setServer(mqtt_server, mqtt_port);
    client.setCallback(callback);
}

void loop() {
    if (!client.connected()) reconnect();
    client.loop();
}