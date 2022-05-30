package auth

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

// Implementation of the repository in this service.
type AuthController struct{}

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Signs in a user and gives them a JWT.
func (h *AuthController) SignInUser(c *fiber.Ctx) error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create a struct so the request body can be mapped here.
	type loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Create a struct for our custom JWT payload.
	type jwtClaims struct {
		UserID string `json:"uid"`
		User   string `json:"user"`
		jwt.StandardClaims
	}

	// Get request body.
	request := &loginRequest{}
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// If both username and password are incorrect, do not allow access.
	if request.Username != os.Getenv("API_USERNAME") || request.Password != os.Getenv("API_PASSWORD") {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Wrong username or password!",
		})
	}

	// Send back JWT as a cookie.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtClaims{
		os.Getenv("API_USERID"),
		os.Getenv("API_USERNAME"),
		jwt.StandardClaims{
			Audience:  "docker-mariadb-clean-arch-users",
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "docker-mariadb-clean-arch",
		},
	})
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    signedToken,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24),
		Secure:   false,
		HTTPOnly: true,
	})

	// Send also response.
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"token":  signedToken,
	})
}

// Logs out user and removes their JWT.
func (h *AuthController) SignOutUser(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "loggedOut",
		Path:     "/",
		Expires:  time.Now().Add(time.Second * 10),
		Secure:   false,
		HTTPOnly: true,
	})

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Logged out successfully!",
	})
}

// A single private route that only logged in users can access.
func (h *AuthController) PrivateRoute(c *fiber.Ctx) error {
	// Give form to our output response.
	type jwtResponse struct {
		UserID interface{} `json:"uid"`
		User   interface{} `json:"user"`
		Iss    interface{} `json:"iss"`
		Aud    interface{} `json:"aud"`
		Exp    interface{} `json:"exp"`
	}

	// Prepare our variables to be displayed.
	jwtData := c.Locals("user").(*jwt.Token)
	claims := jwtData.Claims.(jwt.MapClaims)

	// Shape output response.
	jwtResp := &jwtResponse{
		UserID: claims["uid"],
		User:   claims["user"],
		Iss:    claims["iss"],
		Aud:    claims["aud"],
		Exp:    claims["exp"],
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Welcome to the private route!",
		"jwtData": jwtResp,
	})
}
