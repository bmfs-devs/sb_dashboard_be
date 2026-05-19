package handlers

import (
	resource "github.com/bmfs-devs/sb_dashboard_be/dal/resources"
	"github.com/bmfs-devs/sb_dashboard_be/models"
	mUsecase "github.com/bmfs-devs/sb_dashboard_be/usecases/models"
	"github.com/gofiber/fiber/v2"
)

func OnOffHandler(c *fiber.Ctx) error {
	request := new(models.ACRemoteDTO)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ResponseGeneral{
			Status:       400,
			ErrorMessage: "Invalid request body",
			Data:         nil,
		})
	}

	if request == nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ResponseGeneral{
			Status:       400,
			ErrorMessage: "Invalid request body",
			Data:         nil,
		})
	}

	if len(request.ACStatuses) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.ResponseGeneral{
			Status:       400,
			ErrorMessage: "AC statuses cannot be empty",
			Data:         nil,
		})
	}

	onOffRequest := mUsecase.ACRemoteRequest{
		Ctx: c.Context(),
	}
	for _, acStatus := range request.ACStatuses {
		uACStatus := mUsecase.ACStatus{
			ACNumber:    acStatus.ACNumber,
			Status:      acStatus.Status,
			Temperature: 24, // Default temperature, you can modify this as needed
		}
		onOffRequest.ACStatuses = append(onOffRequest.ACStatuses, uACStatus)
	}
	response, err := resource.Usecase.OnOffAC(
		onOffRequest,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseGeneral{
			Status:       500,
			ErrorMessage: err.Error(),
			Data:         nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(models.ResponseGeneral{
		Status:       200,
		ErrorMessage: "",
		Data:         response,
	})
}

func SetTemperatureHandler(c *fiber.Ctx) error {
	request := new(models.ACTemperatureDTO)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ResponseGeneral{
			Status:       400,
			ErrorMessage: "Invalid request body",
			Data:         nil,
		})
	}
	if request == nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ResponseGeneral{
			Status:       400,
			ErrorMessage: "Invalid request body",
			Data:         nil,
		})

	}

	if len(request.ACTemperatures) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.ResponseGeneral{
			Status:       400,
			ErrorMessage: "AC temperatures cannot be empty",
			Data:         nil,
		})
	}

	temperatureRequest := mUsecase.ACRemoteRequest{
		Ctx: c.Context(),
	}
	for _, acTemperature := range request.ACTemperatures {
		uACStatus := mUsecase.ACStatus{
			ACNumber:    acTemperature.ACNumber,
			Temperature: acTemperature.Temperature,
		}
		temperatureRequest.ACStatuses = append(temperatureRequest.ACStatuses, uACStatus)
	}

	response, err := resource.Usecase.SetTemperature(
		temperatureRequest,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseGeneral{
			Status:       500,
			ErrorMessage: err.Error(),
			Data:         nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(models.ResponseGeneral{
		Status:       200,
		ErrorMessage: "",
		Data:         response,
	})
}

func GetTemperatureHandler(c *fiber.Ctx) error {

	acNumberParam := c.Query("ac_number")

	resp, err := resource.Usecase.GetTemperature(mUsecase.GetTemperatureRequest{
		Ctx:      c.Context(),
		ACNumber: acNumberParam,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseGeneral{
			Status:       500,
			ErrorMessage: err.Error(),
			Data:         nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(models.ResponseGeneral{
		Status:       200,
		ErrorMessage: "",
		Data:         resp,
	})
}
