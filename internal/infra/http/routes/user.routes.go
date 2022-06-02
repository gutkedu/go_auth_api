package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gutkedu/golang_api/internal/infra/http/middlewares"
	"github.com/gutkedu/golang_api/internal/modules/user"
)

func NewUserController(userRoute fiber.Router, us user.UserUseCase) {
	controller := &user.UserController{
		UserUseCase: us,
	}
	userRoute.Get("",
		middlewares.EnsureAuthentication(),
		controller.GetUsersController)
	userRoute.Post("",
		middlewares.EnsureAuthentication(),
		controller.CreateUserController)
	userRoute.Get("/:userID",
		middlewares.EnsureAuthentication(),
		controller.GetUserController)
	userRoute.Put("/:userID",
		middlewares.EnsureAuthentication(),
		controller.CheckIfUserExistsMiddleware,
		controller.UpdateUserController)
	userRoute.Delete("/:userID",
		middlewares.EnsureAuthentication(),
		controller.CheckIfUserExistsMiddleware,
		controller.DeleteUserController)
}
