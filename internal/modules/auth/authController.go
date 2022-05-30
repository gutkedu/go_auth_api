package auth

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Implementation of the repository in this service.
type AuthController struct {
	authUseCase AuthUseCase
}

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Login get user and password
func (h *AuthController) LoginController(c *fiber.Ctx) error {
	var input LoginInput
	//var ud UserData

	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}
	email := input.email
	//pass := input.Password

	user, err := h.authUseCase.findUserByEmail(customContext, email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{"status": "error", "message": "Error on email", "data": err})
	}

	if user == nil {
		fmt.Println("user nil")
	}

	/*
		if user == nil {
				return c.Status(fiber.StatusUnauthorized).JSON(
					fiber.Map{"status": "error", "message": "User not found", "data": err})
			}

		if !CheckPasswordHash(pass, ud.Password) {
			return c.Status(fiber.StatusUnauthorized).JSON(
				fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
		}

		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = ud.Name
		claims["user_id"] = ud.ID
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		tokenString, err := token.SignedString([]byte("be099fcd19a5a65037e9a1f594379027"))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	*/
	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": "tokenString"})
}
