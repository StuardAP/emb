package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/StuardAP/emb/pkg/domain/service"
)

type EmbedHandler struct {
	Service *service.EmbedService
}

func NewEmbedHandler(service *service.EmbedService) *EmbedHandler {
	return &EmbedHandler{
		Service: service,
	}
}

func (h *EmbedHandler) HandlePostEmbed(c *fiber.Ctx) error {
	var req map[string]interface{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse json"})
	}

	text, ok := req["text"].(string)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing or invalid 'text' field"})
	}

	embeddings, err := h.Service.GetEmbeddings(text)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"embeddings": embeddings})
}

func SetupRoutes(app *fiber.App, embedHandler *EmbedHandler) {
	app.Post("/embed", embedHandler.HandlePostEmbed)
}
