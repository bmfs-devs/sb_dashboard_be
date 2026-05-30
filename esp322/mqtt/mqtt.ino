#include <WiFi.h>
#include <WiFiClientSecure.h>
#include <PubSubClient.h>

const char* ssid         = "pts";
const char* password     = "akudankamu";
const int   mqtt_port    = 8883;
const char* mqtt_server = "";
const char* mqtt_user   = "";
const char* mqtt_pass   = "";
const char* AC_TOPIC     = "ac_topic";
const char* SENSOR_TOPIC = "sensor_topic";

WiFiClientSecure espClient;
PubSubClient client(espClient);

void callback(char* topic, byte* payload, unsigned int length) {
    String msg = "";
    for (int i = 0; i < length; i++) msg += (char)payload[i];
    msg.trim();
    Serial.println("[MQTT] Received: " + msg);

    if (String(topic) == AC_TOPIC) handleAC(msg);
}

void setup_wifi() {
    WiFi.disconnect(true);
    WiFi.mode(WIFI_STA);
    WiFi.begin(ssid, password);
    int timeout = 0;
    while (WiFi.status() != WL_CONNECTED) {
        delay(1000);
        if (++timeout > 20) { Serial.println("[WiFi] Timeout."); return; }
    }
    Serial.println("[WiFi] Connected: " + WiFi.localIP().toString());
    espClient.setInsecure();
}

void reconnect() {
    while (!client.connected()) {
        if (WiFi.status() != WL_CONNECTED) setup_wifi();
        String clientId = "ESP32_" + String(random(0xffff), HEX);
        if (client.connect(clientId.c_str(), mqtt_user, mqtt_pass)) {
            Serial.println("[MQTT] Connected");
            client.subscribe(AC_TOPIC);
        } else {
            Serial.printf("[MQTT] Failed (rc=%d). Retry in 5s...\n", client.state());
            delay(5000);
        }
    }
}

void setup() {
    Serial.begin(115200);
    setup_wifi();
    client.setServer(mqtt_server, mqtt_port);
    client.setCallback(callback);
    setupAC();
    setupSensor();
}

void loop() {
    if (!client.connected()) reconnect();
    client.loop();
    publishSensorData();
}