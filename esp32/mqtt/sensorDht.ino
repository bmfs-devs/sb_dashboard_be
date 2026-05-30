#include <DHT.h>

#define DHT_PIN           23
#define DHT_TYPE          DHT22
#define SENSOR_TOPIC      "sensor_topic"
#define SENSOR_INTERVAL   10000  // publish every 10 seconds

DHT dht(DHT_PIN, DHT_TYPE);
unsigned long lastSensorPublish = 0;

void setupSensor() {
    dht.begin();
    Serial.println("[DHT22] Sensor initialized.");
}

void publishSensorData() {
    unsigned long now = millis();
    if (now - lastSensorPublish < SENSOR_INTERVAL) return;
    lastSensorPublish = now;

    float humidity = dht.readHumidity();
    float tempC    = dht.readTemperature();

    if (isnan(humidity) || isnan(tempC)) {
        Serial.println("[DHT22] Failed to read sensor!");
        return;
    }

    String payload = "{\"temp\":"     + String(tempC, 1) +
                     ",\"humidity\":" + String(humidity, 1) + "}";

    client.publish(SENSOR_TOPIC, payload.c_str());
    Serial.println("[DHT22] Published: " + payload);
}