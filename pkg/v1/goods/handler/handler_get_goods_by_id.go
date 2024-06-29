package handler

import (
	coregoods "github.com/11SF/core-family-management/pkg/v1/goods/core"
	"github.com/11SF/core-family-management/validator"
	"github.com/gofiber/fiber/v2"
)

type getGoodsByIdHandler struct {
	getGoodsById coregoods.GetGoodsByIdFunc
}

type getGoodsByIdRequest struct {
	ID string `json:"id" validate:"required"`
}

func NewGetGoodsListHandler(getGoodsById coregoods.GetGoodsByIdFunc) *getGoodsByIdHandler {
	return &getGoodsByIdHandler{
		getGoodsById: getGoodsById,
	}
}

func (h *getGoodsByIdHandler) Handler(c *fiber.Ctx) error {

	ctx := c.Context()

	req := new(getGoodsByIdRequest)
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "BAD_REQUEST",
			"message": err.Error(),
		})
	}

	err = validator.Validate.Struct(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "VALIDATION_ERROR",
			"message": err.Error(),
		})
	}

	goods, err := h.getGoodsById(ctx, req.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "INTERNAL_SERVER_ERROR",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(goods)
}
