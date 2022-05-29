package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gutkedu/golang_api/internal/modules/user"
)

func NewUserController(userRoute fiber.Router, us user.UserUseCase) {
	controller := &user.UserController{
		UserUseCase: us,
	}
	userRoute.Get("", controller.GetUsersController)
	userRoute.Post("", controller.CreateUserController)
	userRoute.Get("/:userID", controller.GetUserController)
	userRoute.Put("/:userID", controller.CheckIfUserExistsMiddleware, controller.UpdateUserController)
	userRoute.Delete("/:userID", controller.CheckIfUserExistsMiddleware, controller.DeleteUserController)
}
