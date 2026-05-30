#include <DHT.h>

#define DHT_PIN          23
#define DHT_TYPE         DHT22
#define SENSOR_INTERVAL  10000

DHT dht(DHT_PIN, DHT_TYPE);
unsigned long lastPublish = 0;

void setupSensor() {
    dht.begin();
    Serial.println("[DHT22] Initialized.");
}

void publishSensorData() {
    if (millis() - lastPublish < SENSOR_INTERVAL) return;
    lastPublish = millis();

    float humidity = dht.readHumidity();
    float tempC    = dht.readTemperature();

    if (isnan(humidity) || isnan(tempC)) {
        Serial.println("[DHT22] Read failed!");
        return;
    }

    String payload = "{\"temp\":"     + String(tempC, 1) +
                     ",\"humidity\":" + String(humidity, 1) + "}";
    client.publish(SENSOR_TOPIC, payload.c_str());
    Serial.println("[DHT22] Published: " + payload);
}