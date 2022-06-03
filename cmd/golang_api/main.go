package main

import (
	"log"

	"github.com/gutkedu/golang_api/internal/infra/http"
)

func main() {
	app := http.Setup()
	log.Fatal(app.Listen(":3333"))
}
