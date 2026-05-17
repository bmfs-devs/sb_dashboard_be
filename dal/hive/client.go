package hive

import (
	"crypto/tls"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	MQTTClient mqtt.Client
)

func InitHiveMQ() {
	mqttPass := os.Getenv("MQTT_PASS")
	mqttuser := os.Getenv("MQTT_USER")
	mqttAddr := os.Getenv("MQTT_ADDR")
	if mqttPass == "" || mqttuser == "" || mqttAddr == "" {
		log.Fatalf("MQTT credentials are not set in environment variables")
	}

	opts := mqtt.NewClientOptions().AddBroker(mqttAddr)
	opts.SetClientID("go_backend_global")
	opts.SetUsername(mqttuser)
	opts.SetPassword(mqttPass)
	opts.SetTLSConfig(&tls.Config{InsecureSkipVerify: false})
	opts.SetCleanSession(true)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Cloud Connection Error: %v", token.Error())
	}

	MQTTClient = client
	log.Println("Successfully connected to HiveMQ Cloud!")
}
