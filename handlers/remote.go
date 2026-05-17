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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if request == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if len(request.ACStatuses) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "AC statuses cannot be empty",
		})
	}

	onOffRequest := mUsecase.ACRemoteRequest{}
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
