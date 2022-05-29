package user

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// If user does not exist, do not allow one to access the API.
func (h *UserController) CheckIfUserExistsMiddleware(c *fiber.Ctx) error {
	// Create a new customized context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Fetch parameter.
	targetedUserID, err := uuid.Parse(c.Params("userID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Please specify a valid user ID!",
		})
	}

	// Check if user exists.
	searchedUser, err := h.UserUseCase.GetUser(customContext, targetedUserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	if searchedUser == nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "There is no user with this ID!",
		})
	}

	// Store in locals for further processing in the real handler.
	c.Locals("userID", targetedUserID)
	return c.Next()
}
