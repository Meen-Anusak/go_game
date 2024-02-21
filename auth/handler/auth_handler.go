package handler

import (
	"github.com/gofiber/fiber/v2"
	"go_game/auth/core/domain"
	"go_game/auth/core/services"
)

type HttpAuthHandler struct {
	service services.AuthService
}

func NewHttpAuthHandler(service *services.AuthServiceImpl) *HttpAuthHandler {
	return &HttpAuthHandler{service: service}
}

// Register godoc
// @Summary Register
// @Description Register
// @Param register body domain.Login true "Register"
// @Tags Auth
// @Accept  application/json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} domain.User
// @Router /auth/register [post]
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

// Login godoc
// @Summary Login
// @Description Login
// @Param login body domain.Login true "Login"
// @Tags Auth
// @Accept  application/json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} domain.Auth
// @Router /auth/login [post]
func (h *HttpAuthHandler) Login(ctx *fiber.Ctx) error {
	var login domain.Login

	if err := ctx.BodyParser(&login); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	user, err := h.service.Login(login)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(user)

}
