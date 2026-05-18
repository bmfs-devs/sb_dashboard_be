package usecases

import (
	"context"
	"fmt"
	"log"

	"github.com/bmfs-devs/sb_dashboard_be/constants"
	mHive "github.com/bmfs-devs/sb_dashboard_be/repositories/queue/hive/models"
	mRedis "github.com/bmfs-devs/sb_dashboard_be/repositories/redis/models"
	mUsecaseRemote "github.com/bmfs-devs/sb_dashboard_be/usecases/models"
	"github.com/bmfs-devs/sb_dashboard_be/utils"
)

func (u *Usecase) OnOffAC(params mUsecaseRemote.ACRemoteRequest) (mUsecaseRemote.ACRemoteResponse, error) {
	// Validate AC numbers and their corresponding hex codes
	for _, acStatus := range params.ACStatuses {
		switch acStatus.ACNumber {
		case 1:
			if !u.validateACHex(acStatus.ACNumber, constants.OnHex) {
				return mUsecaseRemote.ACRemoteResponse{}, fmt.Errorf("invalid AC number %d for on status", acStatus.ACNumber)
			}
		case 0:
			if !u.validateACHex(acStatus.ACNumber, constants.OffHex) {
				return mUsecaseRemote.ACRemoteResponse{}, fmt.Errorf("invalid AC number %d for off status", acStatus.ACNumber)
			}
		}

	}
	response := mUsecaseRemote.ACRemoteResponse{}
	// Publish the corresponding hex codes to the Hive repository
	for _, acStatus := range params.ACStatuses {
		topic := utils.GetACTopic(acStatus.ACNumber)
		if topic == "" {
			return mUsecaseRemote.ACRemoteResponse{}, fmt.Errorf("no topic found for AC number %d", acStatus.ACNumber)
		}

		var mapHex map[int]string
		switch acStatus.ACNumber {
		case 1:
			mapHex = constants.OnHex
		case 0:
			mapHex = constants.OffHex
		}
		hexCode := utils.GetFormatHex(acStatus.ACNumber, mapHex)
		err := u.HiveRepo.Publish(mHive.HiveMessage{
			Topic:   topic,
			Payload: hexCode,
		})
		if err != nil {
			log.Printf("Failed to publish message for AC number %d: %v", acStatus.ACNumber, err)
			return mUsecaseRemote.ACRemoteResponse{}, fmt.Errorf("failed to publish message for AC number %d: %v", acStatus.ACNumber, err)
		}

		// Turn of ac skip default temperature
		if acStatus.Status == 0 {
			response.ACStatuses = append(response.ACStatuses, mUsecaseRemote.ACStatus{
				ACNumber: acStatus.ACNumber,
				Status:   acStatus.Status,
			})
			continue
		}

		// Publish the default temperature hex code to the Hive repository
		tempHexCode := utils.GetFormatHex(acStatus.ACNumber, constants.Temperature24)
		err = u.HiveRepo.Publish(mHive.HiveMessage{
			Topic:   topic,
			Payload: tempHexCode,
		})
		if err != nil {
			log.Printf("Failed to publish default temperature for AC number %d: %v", acStatus.ACNumber, err)
			return mUsecaseRemote.ACRemoteResponse{}, fmt.Errorf("failed to publish message for AC number %d: %v", acStatus.ACNumber, err)
		}

		response.ACStatuses = append(response.ACStatuses, mUsecaseRemote.ACStatus{
			ACNumber:    acStatus.ACNumber,
			Status:      acStatus.Status,
			Temperature: 24, // Default temperature, you can modify this as needed
		})
	}

	defer func(ctx context.Context) {
		if ctx == nil {
			ctx = context.Background()
		}
		for _, acStatus := range params.ACStatuses {
			if acStatus.Status == 1 {
				err := u.RedisRepo.HSet(mRedis.RedisMessageHSetRequest{
					Ctx: ctx,
					Key: utils.RedisKeyACTemperature(),
					Field: map[string]interface{}{
						fmt.Sprintf("%d", acStatus.ACNumber): acStatus.Temperature,
					},
				})
				if err != nil {
					log.Printf("Failed to set temperature for AC number %d: %v", acStatus.ACNumber, err)
				}
				continue
			}
			if acStatus.Status == 0 {
				temperature, err := u.RedisRepo.HGet(mRedis.RedisMessageHGetRequest{
					Ctx:   ctx,
					Key:   utils.RedisKeyACTemperature(),
					Field: fmt.Sprintf("%d", acStatus.ACNumber),
				})
				if err != nil {
					log.Printf("Failed to get temperature for AC number %d: %v", acStatus.ACNumber, err)
					continue
				}

				if temperature != "" {
					err = u.RedisRepo.HDel(mRedis.RedisMessageHGetRequest{
						Ctx:   ctx,
						Key:   utils.RedisKeyACTemperature(),
						Field: fmt.Sprintf("%d", acStatus.ACNumber),
					})
					if err != nil {
						log.Printf("Failed to delete temperature for AC number %d: %v", acStatus.ACNumber, err)
					}
				}
			}
		}
	}(params.Ctx)
	return response, nil

}

func (u *Usecase) validateACHex(acNumber int, maps map[int]string) bool {
	if _, exists := maps[acNumber]; exists {
		return true
	}
	return false
}

func (u *Usecase) SetTemperature(params mUsecaseRemote.ACRemoteRequest) (mUsecaseRemote.ACRemoteResponse, error) {
	// Validate Temperature values and their corresponding hex codes
	for _, acTemp := range params.ACStatuses {
		if acTemp.Temperature < constants.TemperatureMin || acTemp.Temperature > constants.TemperatureMax {
			return mUsecaseRemote.ACRemoteResponse{}, fmt.Errorf("invalid temperature %d for AC number %d", acTemp.Temperature, acTemp.ACNumber)
		}

		temperatureHex, exists := constants.TemperatureHex[acTemp.Temperature]
		if !exists {
			return mUsecaseRemote.ACRemoteResponse{}, fmt.Errorf("no hex code found for temperature %d", acTemp.Temperature)
		}
		if !u.validateACHex(acTemp.ACNumber, temperatureHex) {
			return mUsecaseRemote.ACRemoteResponse{}, fmt.Errorf("invalid AC number %d for temperature %d", acTemp.ACNumber, acTemp.Temperature)
		}
	}

	res := mUsecaseRemote.ACRemoteResponse{}
	// Publish the corresponding hex codes to the Hive repository
	for _, acTemp := range params.ACStatuses {
		// Check redis for current AC status
		currentStatus, err := u.RedisRepo.HGet(mRedis.RedisMessageHGetRequest{
			Ctx:   params.Ctx,
			Key:   utils.RedisKeyACTemperature(),
			Field: fmt.Sprintf("%d", acTemp.ACNumber),
		})
		if err != nil || currentStatus == "" {
			log.Printf("Failed to get temperature for AC number %d: %v, or AC is not ON, skipping...", acTemp.ACNumber, err)
			continue
		}

		// If AC is ON, proceed to publish the new temperature
		topic := utils.GetACTopic(acTemp.ACNumber)
		if topic == "" {
			return mUsecaseRemote.ACRemoteResponse{}, fmt.Errorf("no topic found for AC number %d", acTemp.ACNumber)
		}

		temperatureHex := utils.GetFormatHex(acTemp.ACNumber, constants.TemperatureHex[acTemp.Temperature])
		err = u.HiveRepo.Publish(mHive.HiveMessage{
			Topic:   topic,
			Payload: temperatureHex,
		})
		if err != nil {
			log.Printf("Failed to publish temperature for AC number %d: %v", acTemp.ACNumber, err)
			return mUsecaseRemote.ACRemoteResponse{}, fmt.Errorf("failed to publish message for AC number %d: %v", acTemp.ACNumber, err)
		}

		// Update Redis with the new temperature
		err = u.RedisRepo.HSet(mRedis.RedisMessageHSetRequest{
			Ctx: params.Ctx,
			Key: utils.RedisKeyACTemperature(),
			Field: map[string]interface{}{
				fmt.Sprintf("%d", acTemp.ACNumber): acTemp.Temperature,
			},
		})
		if err != nil {
			log.Printf("Failed to set temperature for AC number %d: %v", acTemp.ACNumber, err)
		}

		res.ACStatuses = append(res.ACStatuses, mUsecaseRemote.ACStatus{
			ACNumber:    acTemp.ACNumber,
			Status:      1, // Assuming the AC is ON if we're setting the temperature
			Temperature: acTemp.Temperature,
		})
	}

	return res, nil
}
