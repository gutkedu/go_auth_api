package http

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gutkedu/golang_api/internal/infra/gorm"
	"github.com/gutkedu/golang_api/internal/infra/http/routes"
	"github.com/gutkedu/golang_api/internal/modules/user"
	"github.com/gutkedu/golang_api/internal/modules/user/infra/gorm/entities"
)

func Run() {

	pgdb, err := gorm.ConnectToPgDB()
	if err != nil {
		log.Fatal("database connection error: ", err)
	}

	pgdb.AutoMigrate(&entities.User{})

	app := fiber.New(fiber.Config{
		AppName:      "golangAPI",
		ServerHeader: "FiberServer",
	})

	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(etag.New())
	app.Use(favicon.New())
	app.Use(limiter.New(limiter.Config{
		Max: 100,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
				"status":  "fail",
				"message": "You have requested too many in a single time-frame! Please wait another minute!",
			})
		},
	}))
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())

	// Create repositories.
	userRepository := user.NewUserRepository(pgdb)

	// Create all of our services.
	userUseCase := user.NewUserUseCase(userRepository)

	// Prepare our endpoints for the API.
	routes.NewUserController(app.Group("/api/v1/users"), userUseCase)

	// Prepare an endpoint for 'Not Found'.
	app.All("*", func(c *fiber.Ctx) error {
		errorMessage := fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL())
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"status":  "fail",
			"message": errorMessage,
		})
	})

	log.Fatal(app.Listen(":3333"))
}
