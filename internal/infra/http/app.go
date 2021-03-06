package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gutkedu/golang_api/internal/infra/gorm"
	"github.com/gutkedu/golang_api/internal/infra/gorm/migrations"
	"github.com/gutkedu/golang_api/internal/infra/http/middlewares"
	"github.com/gutkedu/golang_api/internal/infra/http/routes"
)

func Setup() *fiber.App {

	//Database
	pgdb, err := gorm.ConnectToPgDB()
	if err != nil {
		log.Fatal("database connection error: ", err)
	}

	//Migrations
	migrations.RunGormMigrations(pgdb)

	app := fiber.New(fiber.Config{
		AppName:      "golangAPI",
		ServerHeader: "FiberServer",
	})

	//Middlewares
	middlewares.FiberMiddleware(app)

	//Router
	routes.RegisterRoutes(app, pgdb)

	return app
}
