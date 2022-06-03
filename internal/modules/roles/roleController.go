package roles

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type RoleController struct {
	RoleUseCase RoleUseCase
}

func (h *RoleController) CreateRoleController(c *fiber.Ctx) error {
	role := &Role{}

	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Parse request body.
	if err := c.BodyParser(role); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if err := role.CreateRoleValidation(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	err := h.RoleUseCase.CreateRole(customContext, role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "Role has been created successfully!",
	})
}

func (h *RoleController) GetRoleController(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	roleID, err := uuid.Parse(c.Params("roleID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Please specify a valid user ID!",
		})
	}

	role, err := h.RoleUseCase.GetRole(customContext, roleID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   role,
	})
}

func (h *RoleController) GetRolesController(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	roles, err := h.RoleUseCase.GetRoles(customContext)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   roles,
	})
}

func (h *RoleController) UpdateRoleController(c *fiber.Ctx) error {
	role := &Role{}

	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Fetch parameter.
	roleID, err := uuid.Parse(c.Params("roleID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Please specify a valid role ID!",
		})
	}

	if err := c.BodyParser(role); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if err := role.CreateRoleValidation(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	err = h.RoleUseCase.UpdateRole(customContext, roleID, role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Role has been updated successfully!",
	})
}

func (h *RoleController) DeleteRoleController(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Fetch parameter.
	roleID, err := uuid.Parse(c.Params("roleID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Please specify a valid role ID!",
		})
	}

	err = h.RoleUseCase.DeleteRole(customContext, roleID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Role has been deleted successfully!",
	})
}
