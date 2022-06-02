package auth

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

// Implementation of the repository in this service.
type AuthController struct {
	AuthUserUseCase AuthUserUseCase
}

func (h *AuthController) GetNewAccessToken(c *fiber.Ctx) error {
	var body AuthRequest

	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}

	if err := body.ValidateLoginInput(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}

	res, err := h.AuthUserUseCase.Execute(customContext, body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   res,
	})
}
