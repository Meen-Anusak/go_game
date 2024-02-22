package handler

import (
	"github.com/gofiber/fiber/v2"
	"go_game/crm/core/domain"
	"go_game/crm/core/services"
)

type HttpBrandHandler struct {
	service services.BrandService
}

func NewHttpBrandHandler(service *services.BrandServiceImpl) *HttpBrandHandler {
	return &HttpBrandHandler{service: service}
}

// GetAllBrand godoc
// @Summary GetAllBrand
// @Description GetAllBrand
// @Tags Brand
// @Accept  application/json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} domain.Brand
// @Router /brand/list [get]
func (h *HttpBrandHandler) GetAllBrand(ctx *fiber.Ctx) error {
	brands, err := h.service.GetAllBrand()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(brands)
}

// CreateNewBrand godoc
// @Summary CreateNewBrand
// @Description CreateNewBrand
// @Param brand body domain.BrandRequestSchema true "NewBrand"
// @Tags Brand
// @Accept  application/json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} domain.BrandResponseSchema
// @Router /brand [post]
func (h *HttpBrandHandler) CreateNewBrand(ctx *fiber.Ctx) error {
	var brand domain.Brand
	if err := ctx.BodyParser(&brand); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	newBrand, err := h.service.CreateNewBrand(brand)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(newBrand)
}
