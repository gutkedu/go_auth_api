package roles

import "github.com/gofiber/fiber/v2"

type RoleController struct {
	RoleUseCase RoleUseCase
}

func (h *RoleController) GetRoleController(c *fiber.Ctx) error {
	return nil
}

func (h *RoleController) GetRolesController(c *fiber.Ctx) error {
	return nil
}

func (h *RoleController) CreateRoleController(c *fiber.Ctx) error {
	return nil
}

func (h *RoleController) UpdateRoleController(c *fiber.Ctx) error {
	return nil
}

func (h *RoleController) DeleteRoleController(c *fiber.Ctx) error {
	return nil
}
