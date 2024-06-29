package handler

import (
	"context"
	"log/slog"

	"github.com/11SF/core-family-management/pkg/v1/datamodel"
	corefamily "github.com/11SF/core-family-management/pkg/v1/family/core"
	"github.com/11SF/core-family-management/validator"
	"github.com/11SF/go-common/response"
	"github.com/gofiber/fiber/v2"
)

type getFamilyInfoHandler struct {
	getFamilyInfo corefamily.GetFamilyInfo
}

type getFamilyInfoRequest struct {
	FamilyId string `json:"familyId" validate:"required,uuid"`
}

type getFamilyInfoResponse *datamodel.Family

func NewGetFamilyInfoHandler(getFamilyInfo corefamily.GetFamilyInfo) *getFamilyInfoHandler {
	return &getFamilyInfoHandler{getFamilyInfo: getFamilyInfo}
}

func (h *getFamilyInfoHandler) Handler(c *fiber.Ctx) error {

	ctx := context.WithValue(c.Context(), "userId", "test")

	req := new(getFamilyInfoRequest)

	err := c.ParamsParser(req)
	if err != nil {
		slog.Error("failed to parse request", slog.String("error", err.Error()), slog.String("tag", "get family info handler"))
		return response.NewFiberResponseError(c, fiber.StatusOK, response.NewError("CFM400", err.Error()))
	}

	err = validator.Validate.Struct(req)
	if err != nil {
		slog.Error("failed to validate request", slog.String("error", err.Error()), slog.String("tag", "get family info handler"))
		return response.NewFiberResponseError(c, fiber.StatusOK, response.NewError("CFM401", err.Error()))
	}

	res, err := h.getFamilyInfo(ctx, req.FamilyId)
	if err != nil {
		slog.Error("failed to get family info", slog.String("error", err.Error()), slog.String("tag", "get family info handler"))
		return response.NewFiberResponseError(c, fiber.StatusOK, err)
	}

	resp := res

	return response.NewFiberResponse(c, response.SuccessCode, resp)
}
