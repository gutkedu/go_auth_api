package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gutkedu/golang_api/internal/modules/roles"
)

func NewRolesController(
	rolesRoute fiber.Router,
	uc roles.RoleUseCase) {
	controller := &roles.RoleController{
		RoleUseCase: uc,
	}

	rolesRoute.Post("", controller.CreateRoleController)
}
