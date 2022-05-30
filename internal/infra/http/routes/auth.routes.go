package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gutkedu/golang_api/internal/modules/auth"
)

// Creates a new authentication handler.
func NewAuthController(authRoute fiber.Router, us auth.AuthUseCase) {
	controller := &auth.AuthController{}
	// Declare routing for specific routes.
	authRoute.Post("/login", controller.LoginController)
}
