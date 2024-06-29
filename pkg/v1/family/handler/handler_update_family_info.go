package handler

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/11SF/core-family-management/pkg/v1/datamodel"
	corefamily "github.com/11SF/core-family-management/pkg/v1/family/core"
	commonValidator "github.com/11SF/core-family-management/validator"
	"github.com/11SF/go-common/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type updateFamilyInfoHandler struct {
	updateFamilyInfo corefamily.UpdateFamilyInfo
}

type updateFamilyInfoRequest struct {
	FamilyId        string   `json:"familyId" validate:"required,uuid"`
	Name            string   `json:"name" validate:"required,max=50"`
	Platform        string   `json:"platform" validate:"required"`
	DueDate         int      `json:"dueDate" validate:"required,gte=1,lte=31"`
	PromptPayNumber string   `json:"promptPayNumber" validate:"required"`
	Prices          []prices `json:"prices" validate:"omitempty,dive"`
}

func NewUpdateFamilyInfoHandler(updateFamilyInfo corefamily.UpdateFamilyInfo) *updateFamilyInfoHandler {
	return &updateFamilyInfoHandler{updateFamilyInfo: updateFamilyInfo}
}

func (h *updateFamilyInfoHandler) Handler(c *fiber.Ctx) error {

	ctx := context.WithValue(c.Context(), "userId", "test")
	req := new(updateFamilyInfoRequest)
	err := c.BodyParser(req)
	if err != nil {
		slog.Error("fail to parse request", slog.String("error", err.Error()), slog.String("tag", "update family info handler"))
		return response.NewFiberResponseError(c, http.StatusOK, response.NewError("CFM400", err.Error()))
	}

	err = commonValidator.Validate.Struct(req)
	if err != nil {
		if err, ok := err.(validator.ValidationErrors); ok {
			slog.Error("fail to validate request", slog.String("error", err.Error()), slog.String("tag", "update family info handler"))
			return response.NewFiberResponseError(c, http.StatusOK, response.NewError("CFM401", err.Error()))
		}
		slog.Error("fail to validate request", slog.String("error", err.Error()), slog.String("tag", "update family info handler"))
		return response.NewFiberResponseError(c, http.StatusOK, response.NewError("CFM401", err.Error()))
	}

	pricesPayload := make([]datamodel.Prices, len(req.Prices))
	for i, price := range req.Prices {
		pricesPayload[i] = datamodel.Prices{
			Price: price.Price,
			Month: price.Month,
		}
	}

	payload := &datamodel.Family{
		ID:              req.FamilyId,
		Name:            req.Name,
		Platform:        req.Platform,
		DueDate:         req.DueDate,
		PromptPayNumber: req.PromptPayNumber,
		Prices:          &pricesPayload,
	}

	err = h.updateFamilyInfo(ctx, payload)
	if err != nil {
		slog.Error("fail to update family info", slog.String("error", err.Error()), slog.String("tag", "update family info handler"))
		return response.NewFiberResponseError(c, http.StatusOK, response.NewError("CFM402", err.Error()))
	}

	return response.NewFiberResponse(c, response.SuccessCode, nil)
}
