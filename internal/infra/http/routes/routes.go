package routes

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gutkedu/golang_api/internal/modules/auth"
	"github.com/gutkedu/golang_api/internal/modules/roles"
	"github.com/gutkedu/golang_api/internal/modules/user"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, pgdb *gorm.DB) {
	// Create repositories.
	userRepository := user.NewUserRepository(pgdb)
	rolesRepository := roles.NewRoleRepository(pgdb)

	// Create all of our services.
	userUseCase := user.NewUserUseCase(userRepository, rolesRepository)
	authUserUseCase := auth.NewAuthUserUseCase(userRepository)
	roleUseCase := roles.NewRoleUseCase(rolesRepository)

	// Prepare our endpoints for the API.
	NewAuthController(app.Group("/api/v1/auth"), authUserUseCase)
	NewUserController(app.Group("/api/v1/users"), userUseCase)
	NewRolesController(app.Group("/api/v1/roles"), roleUseCase)

	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	app.Get("/routes", func(c *fiber.Ctx) error {
		routesData, err := json.Marshal(app.Stack())
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status":  "fail",
				"message": err,
			})
		}
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status": "success",
			"data":   string(routesData),
		})
	})

	// Prepare an endpoint for 'Not Found'.
	app.All("*", func(c *fiber.Ctx) error {
		errorMessage := fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL())
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"status":  "fail",
			"message": errorMessage,
		})
	})
}
