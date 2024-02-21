package handler

import (
	"go_game/game/core/domain"
	service "go_game/game/core/services"

	"github.com/gofiber/fiber/v2"
)

type HttpPlayerHandler struct {
	service service.PlayerService
}

func NewHttpPlayerHandler(service *service.PlayerServiceImpl) *HttpPlayerHandler {
	return &HttpPlayerHandler{service: service}
}

// GetAllPlayer godoc
// @Summary GetAllPlayer
// @Description GetAllPlayer
// @Tags Player
// @Accept  application/json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} domain.Player
// @Router /player/list [get]
func (h *HttpPlayerHandler) GetAllPlayer(ctx *fiber.Ctx) error {
	player, err := h.service.GetAllPlayer()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(player)

}

// CreateNewPlayer godoc
// @Summary CreateNewPlayer
// @Description CreateNewPlayer
// @Param player body domain.Player true "NewPlayer"
// @Tags Player
// @Accept  application/json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} domain.Player
// @Router /player [post]
func (h *HttpPlayerHandler) CreateNewPlayer(ctx *fiber.Ctx) error {
	var player domain.Player
	if err := ctx.BodyParser(&player); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	if err := h.service.CreateNewPlayer(player); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Create New Player Success"})
}

// GetPlayerById godoc
// @Summary GetPlayerById
// @Description GetPlayerById
// @Param player_id path string true "Player ID"
// @Tags Player
// @Accept  application/json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} domain.Player
// @Router /player [get]
func (h *HttpPlayerHandler) GetPlayerById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	player, err := h.service.GetPlayerById(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(player)
}

// UpdatePlayer godoc
// @Summary UpdatePlayer
// @Description UpdatePlayer
// @Param player body domain.Player true "UpdatePlayer"
// @Tags Player
// @Accept  application/json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} domain.Player
// @Router /player [patch]
func (h *HttpPlayerHandler) UpdatePlayer(ctx *fiber.Ctx) error {
	var player domain.Player
	if err := ctx.BodyParser(&player); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	updatePlayer, err := h.service.UpdatePlayer(player)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(updatePlayer)
}
