package adapter

import (
	"fmt"
	"go_game/game/core"

	"github.com/gofiber/fiber/v2"
)

type HttpPlayerHandler struct {
	service core.PlayerService
}

func NewHttpPlayerHandler(service core.PlayerService) *HttpPlayerHandler {
	return &HttpPlayerHandler{service: service}
}

func (h *HttpPlayerHandler) CreateNewPlayerHandler(ctx *fiber.Ctx) error {
	var player core.Player
	if err := ctx.BodyParser(&player); err != nil {
		fmt.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := h.service.CreateNewPlayer(player); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(player)

}
