package handler

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/11SF/core-family-management/pkg/v1/datamodel"
	corefamily "github.com/11SF/core-family-management/pkg/v1/family/core"
	"github.com/11SF/core-family-management/validator"
	"github.com/11SF/go-common/response"
	"github.com/gofiber/fiber/v2"
)

type createFamilyHandler struct {
	createFamily corefamily.CreateFamilyFunc
}

type createFamilyRequest struct {
	Name            string   `json:"name" validate:"required"`
	Platform        string   `json:"platform" validate:"required"`
	DueDate         int      `json:"dueDate" validate:"required"`
	PromptPayNumber string   `json:"promptPayNumber" validate:"required"`
	Prices          []prices `json:"prices" validate:"omitempty,dive"`
}

type prices struct {
	Price float64 `json:"price" validate:"required"`
	Month int     `json:"month" validate:"required"`
}

type createFamilyResponse struct {
	FamilyId string `json:"familyId"`
}

func NewCreateFamilyHandler(createFamily corefamily.CreateFamilyFunc) *createFamilyHandler {
	return &createFamilyHandler{createFamily}
}

func (h createFamilyHandler) Handler(c *fiber.Ctx) error {

	ctx := context.WithValue(c.Context(), "userId", "test")

	req := new(createFamilyRequest)

	err := c.BodyParser(req)
	if err != nil {
		slog.Error("failed to parse request", slog.String("error", err.Error()), slog.String("tag", "create family handler"))
		return response.NewFiberResponseError(c, http.StatusOK, response.NewError("CFM400", err.Error()))
	}

	err = validator.Validate.Struct(req)
	if err != nil {
		slog.Error("failed to validate request", slog.String("error", err.Error()), slog.String("tag", "create family handler"))
		return response.NewFiberResponseError(c, http.StatusOK, response.NewError("CFM401", err.Error()))
	}

	pricesPayload := make([]datamodel.Prices, len(req.Prices))
	for i, price := range req.Prices {
		pricesPayload[i] = datamodel.Prices{
			Price: price.Price,
			Month: price.Month,
		}
	}

	familyPayload := &datamodel.Family{
		Name:            req.Name,
		Platform:        req.Platform,
		DueDate:         req.DueDate,
		PromptPayNumber: req.PromptPayNumber,
		Prices:          &pricesPayload,
	}

	familyId, err := h.createFamily(ctx, familyPayload)
	if err != nil {
		slog.Error("failed to create family", slog.String("error", err.Error()), slog.String("tag", "create family handler"))
		return response.NewFiberResponseError(c, http.StatusOK, err)
	}

	resp := createFamilyResponse{
		FamilyId: familyId,
	}

	return response.NewFiberResponse(c, response.SuccessCode, resp)
}
