package handler

import (
	"github.com/gofiber/fiber/v2"
	"go_game/auth/core/domain"
	"go_game/auth/core/service"
)

type HttpAuthHandler struct {
	service service.AuthService
}

func NewHttpAuthHandler(service *service.AuthServiceImpl) *HttpAuthHandler {
	return &HttpAuthHandler{service: service}
}

func (h *HttpAuthHandler) Register(ctx *fiber.Ctx) error {
	var newUser domain.User

	if err := ctx.BodyParser(&newUser); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := h.service.Register(newUser); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Create New User Success"})
}

func (h *HttpAuthHandler) Login(ctx *fiber.Ctx) error {
	panic("Login Handler")
}
