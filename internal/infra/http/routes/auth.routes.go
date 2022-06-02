package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gutkedu/golang_api/internal/modules/auth"
)

// Creates a new authentication handler.
func NewAuthController(authRoute fiber.Router, au auth.AuthUserUseCase) {
	controller := &auth.AuthController{
		AuthUserUseCase: au,
	}

	// Declare routing for specific routes.
	authRoute.Post("/login", controller.GetNewAccessToken)
	//authRoute.Post("/logout", controller.SignOutUser)
	//authRoute.Get("/private", controller.JWTMiddleware(), controller.PrivateRoute)
}
