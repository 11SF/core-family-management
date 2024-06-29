package handler

import (
	"context"
	"log/slog"
	"net/http"

	corefamily "github.com/11SF/core-family-management/pkg/v1/family/core"
	commonValidator "github.com/11SF/core-family-management/validator"
	"github.com/11SF/go-common/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type deleteFamilyHandler struct {
	deleteFamily corefamily.DeleteFamily
}

type deleteFamilyRequest struct {
	FamilyId string `json:"familyId" validate:"required,uuid"`
}

func NewDeleteFamilyHandler(deleteFamily corefamily.DeleteFamily) *deleteFamilyHandler {
	return &deleteFamilyHandler{deleteFamily: deleteFamily}
}

func (h *deleteFamilyHandler) Handler(c *fiber.Ctx) error {

	ctx := context.WithValue(c.Context(), "userId", "test")

	req := new(deleteFamilyRequest)
	err := c.ParamsParser(req)
	if err != nil {
		return response.NewFiberResponseError(c, http.StatusOK, response.NewError("CFM400", err.Error()))
	}

	err = commonValidator.Validate.Struct(req)
	if err != nil {
		if err, ok := err.(validator.ValidationErrors); ok {
			slog.Error("fail to validate request", slog.String("error", err.Error()), slog.String("tag", "delete family handler"))
			return response.NewFiberResponseError(c, http.StatusOK, response.NewError("CFM401", err.Error()))
		}
		slog.Error("fail to validate request", slog.String("error", err.Error()), slog.String("tag", "delete family handler"))
		return response.NewFiberResponseError(c, http.StatusOK, response.NewError("CFM401", err.Error()))
	}

	err = h.deleteFamily(ctx, req.FamilyId)
	if err != nil {
		slog.Error("fail to delete family", slog.String("error", err.Error()), slog.String("tag", "delete family handler"))
		return response.NewFiberResponseError(c, http.StatusOK, response.NewError("CFM402", err.Error()))
	}

	return response.NewFiberResponse(c, response.SuccessCode, nil)
}
