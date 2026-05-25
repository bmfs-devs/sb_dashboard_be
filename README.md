# sb_dashboard_be

## Setup

### Backend
- Clone The Repository
- Run command `go mod tidy`
- Install Redis server (if use Windows then you must install Docker or WSL)
- Ask the admin regarding credential of HiveMQ
    ```
    // mqttPass := os.Getenv("MQTT_PASS")
	// mqttuser := os.Getenv("MQTT_USER")
	// mqttAddr := os.Getenv("MQTT_ADDR")
    ```
- Run command `go run app/http/http.go`

### Hardware
#### Transmitter and Receiver
- Connect ESP32 to PC/Laptop
- Flash the firmware to ESP32 (`mqtt.ino` and `receiver.ino`)

### Documentation
- Import `sb_dashboard.postman_collection.json` to Postman
- Test the API endpoints
- If you update the API endpoints, please update the Postman collection as well