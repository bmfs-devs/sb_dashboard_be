package usecases

import (
	"fmt"

	"github.com/bmfs-devs/sb_dashboard_be/constants"
	mHive "github.com/bmfs-devs/sb_dashboard_be/repositories/queue/hive/models"
	mUsecaseRemote "github.com/bmfs-devs/sb_dashboard_be/usecases/models"
	"github.com/bmfs-devs/sb_dashboard_be/utils"
)

func (u *Usecase) OnOffAC(params mUsecaseRemote.ACRemoteRequest) (mUsecaseRemote.ACRemoteResponse, error) {
	// Validate AC numbers and their corresponding hex codes
	for _, acStatus := range params.ACStatuses {
		switch acStatus.ACNumber {
		case 1:
			if !u.validateACUnit(acStatus.ACNumber, constants.OnHex) {
				return mUsecaseRemote.ACRemoteResponse{}, fmt.Errorf("invalid AC number %d for on status", acStatus.ACNumber)
			}
		case 0:
			if !u.validateACUnit(acStatus.ACNumber, constants.OffHex) {
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
			return mUsecaseRemote.ACRemoteResponse{}, fmt.Errorf("failed to publish message for AC number %d: %v", acStatus.ACNumber, err)
		}
		response.ACStatuses = append(response.ACStatuses, mUsecaseRemote.ACStatus{
			ACNumber: acStatus.ACNumber,
			Status:   acStatus.Status,
		})
	}
	return response, nil

}

func (u *Usecase) validateACUnit(acNumber int, maps map[int]string) bool {
	if _, exists := maps[acNumber]; exists {
		return true
	}
	return false
}
