package handler

import (
	"github.com/gofiber/fiber/v2"
	"go_game/game/core"
)

type HttpPlayerHandler struct {
	service core.PlayerService
}

func NewHttpPlayerHandler(service *core.PlayerServiceImpl) *HttpPlayerHandler {
	return &HttpPlayerHandler{service: service}
}

func (h *HttpPlayerHandler) GetAllPlayer(ctx *fiber.Ctx) error {
	player, err := h.service.GetAllPlayer()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(player)

}

func (h *HttpPlayerHandler) CreateNewPlayer(ctx *fiber.Ctx) error {
	var player core.Player
	if err := ctx.BodyParser(&player); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	if err := h.service.CreateNewPlayer(player); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Create New Player Success"})
}
