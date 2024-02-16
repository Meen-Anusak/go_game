package handler

import (
	"github.com/gofiber/fiber/v2"
	"go_game/crm/core/domain"
	"go_game/crm/core/services"
)

type HttpUserHandler struct {
	service services.UserService
}

func NewHttpUserHandler(service *services.UserServiceImpl) *HttpUserHandler {
	return &HttpUserHandler{service: service}
}

func (h *HttpUserHandler) CreateNewUser(ctx *fiber.Ctx) error {
	var user domain.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.service.CreateNewUser(user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Create New User Success"})
}
