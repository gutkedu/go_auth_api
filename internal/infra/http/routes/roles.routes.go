package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gutkedu/golang_api/internal/modules/roles"
)

func NewRolesController(
	rolesRoutes fiber.Router,
	uc roles.RoleUseCase) {
	controller := &roles.RoleController{
		RoleUseCase: uc,
	}

	rolesRoutes.Post("", controller.CreateRoleController)
	rolesRoutes.Get("", controller.GetRolesController)
	rolesRoutes.Get("/:roleID", controller.GetRoleController)
	rolesRoutes.Put("/:roleID", controller.UpdateRoleController)
	rolesRoutes.Delete("/:roleID", controller.DeleteRoleController)
}
