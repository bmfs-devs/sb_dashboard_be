package utils

import (
	"fmt"

	"github.com/bmfs-devs/sb_dashboard_be/constants"
)

func GetACTopic(acNumber int) string {
	for topic, acNumbers := range constants.Topics {
		for _, num := range acNumbers {
			if num == acNumber {
				return topic
			}
		}
	}

	return ""
}

func GetFormatHex(acNumber int, mapHex map[int]string) string {
	if hex, exists := mapHex[acNumber]; exists {
		return fmt.Sprintf("%d:%s", acNumber, hex)
	}
	return ""
}

func RedisKeyACTemperature() string {
	return "ac_temperatures"
}
