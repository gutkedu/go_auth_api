package user

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/gutkedu/golang_api/internal/modules/user/infra/gorm/entities"
	"golang.org/x/crypto/bcrypt"
)

// Represents our handler with our use-case / service.
type UserHandler struct {
	userService entities.UserService
}

func NewUserHandler(userRoute fiber.Router, us entities.UserService) {
	handler := &UserHandler{
		userService: us,
	}
	userRoute.Get("", handler.getUsers)
	userRoute.Post("", handler.createUser)

	userRoute.Get("/:userID", handler.getUser)
	userRoute.Put("/:userID", handler.checkIfUserExistsMiddleware, handler.updateUser)
	userRoute.Delete("/:userID", handler.checkIfUserExistsMiddleware, handler.deleteUser)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (h *UserHandler) getUsers(c *fiber.Ctx) error {
	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Get all users.
	users, err := h.userService.GetUsers(customContext)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Return results.
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   users,
	})
}

func (h *UserHandler) getUser(c *fiber.Ctx) error {
	// Create cancellable context.
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

	// Get one user.
	user, err := h.userService.GetUser(customContext, targetedUserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Return results.
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   user,
	})
}

func (h *UserHandler) createUser(c *fiber.Ctx) error {
	user := &entities.User{}

	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Parse request body.
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})
	}
	user.Password = hash

	// Create one user.
	err = h.userService.CreateUser(customContext, user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Return result.
	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "User has been created successfully!",
	})
}

func (h *UserHandler) updateUser(c *fiber.Ctx) error {
	// Initialize variables.
	user := &entities.User{}
	targetedUserID := c.Locals("userID").(uuid.UUID)

	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Parse request body.
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Update one user.
	err := h.userService.UpdateUser(customContext, targetedUserID, user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Return result.
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "User has been updated successfully!",
	})
}

func (h *UserHandler) deleteUser(c *fiber.Ctx) error {
	// Initialize previous user ID.
	targetedUserID := c.Locals("userID").(uuid.UUID)

	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Delete one user.
	err := h.userService.DeleteUser(customContext, targetedUserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Return 204 No Content.
	return c.SendStatus(fiber.StatusNoContent)
}
