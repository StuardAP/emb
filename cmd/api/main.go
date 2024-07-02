package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/StuardAP/emb/internal/delivery/http"
	"github.com/StuardAP/emb/pkg/domain/service"
	http_client "github.com/StuardAP/emb/pkg/infrastructure/http"
)

func main() {
	app := fiber.New()

	httpClient := http_client.NewHTTPClient("localhost:11434")
	embedService := service.NewEmbedService(httpClient)
	embedHandler := http.NewEmbedHandler(embedService)

	http.SetupRoutes(app, embedHandler)

	log.Fatal(app.Listen(":8080"))
}
