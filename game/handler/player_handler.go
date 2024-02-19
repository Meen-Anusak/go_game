package handler

import (
	"go_game/game/core/domain"
	core "go_game/game/core/services"

	"github.com/gofiber/fiber/v2"
)

type HttpPlayerHandler struct {
	service core.PlayerService
}

func NewHttpPlayerHandler(service *core.PlayerServiceImpl) *HttpPlayerHandler {
	return &HttpPlayerHandler{service: service}
}

// GetAllPlayer Handler functions
// GetAllPlayer godoc
// @Summary Get all player
// @Description Get details of all player
// @Tags Player
// @Accept  application/json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} domain.Player
// @Router /player [get]
func (h *HttpPlayerHandler) GetAllPlayer(ctx *fiber.Ctx) error {
	player, err := h.service.GetAllPlayer()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(player)

}

// CreateNewPlayer Handler functions
// CreateNewPlayer godoc
// @Description Create a new player
// @Summary create a new player
// @Tags Player
// @Accept application/json
// @Produce json
// @Param PlayerData body playerInterface true "Player Data"
// @Success 200 {object} domain.Player
// @Security ApiKeyAuth
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
