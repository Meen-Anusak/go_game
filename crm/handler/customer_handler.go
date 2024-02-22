package handler

import (
	"github.com/gofiber/fiber/v2"
	"go_game/crm/core/domain"
	"go_game/crm/core/services"
	"go_game/helper"
)

type HttpCustomerHandler struct {
	service services.CustomerService
}

func NewHttpCustomerHandler(service *services.CustomerServiceImpl) *HttpCustomerHandler {
	return &HttpCustomerHandler{service: service}
}

// GetAllCustomer godoc
// @Summary GetAllCustomer
// @Description GetAllCustomer
// @Tags Customer
// @Accept  application/json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} domain.Customer
// @Router /customer [get]
func (h *HttpCustomerHandler) GetAllCustomer(ctx *fiber.Ctx) error {
	customers, err := h.service.GetAllCustomer()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(customers)
}

// CreateNewCustomer godoc
// @Summary CreateNewCustomer
// @Description CreateNewCustomer
// @Param brand body domain.CustomerRequestSchema true "NewCustomer"
// @Tags Customer
// @Accept  application/json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} domain.CustomerResponseSchema
// @Router /customer [post]
func (h *HttpCustomerHandler) CreateNewCustomer(ctx *fiber.Ctx) error {
	var customer domain.Customer
	if err := ctx.BodyParser(&customer); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := helper.Validate(customer); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":         "V-422",
			"code_message": "validation error",
			"data":         err,
		})
	}
	newCustomer, err := h.service.CreateNewCustomer(customer)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(newCustomer)
}
