package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gutkedu/golang_api/internal/modules/auth"
	"github.com/gutkedu/golang_api/internal/modules/user"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, pgdb *gorm.DB) {
	// Create repositories.
	userRepository := user.NewUserRepository(pgdb)

	// Create all of our services.
	userUseCase := user.NewUserUseCase(userRepository)
	authUserUseCase := auth.NewAuthUserUseCase(userRepository)

	// Prepare our endpoints for the API.
	NewAuthController(app.Group("/api/v1/auth"), authUserUseCase)
	NewUserController(app.Group("/api/v1/users"), userUseCase)

	// Prepare an endpoint for 'Not Found'.
	app.All("*", func(c *fiber.Ctx) error {
		errorMessage := fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL())
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"status":  "fail",
			"message": errorMessage,
		})
	})
}
