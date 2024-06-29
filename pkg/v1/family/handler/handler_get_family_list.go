package handler

import (
	"context"
	"log/slog"

	"github.com/11SF/core-family-management/pkg/v1/datamodel"
	corefamily "github.com/11SF/core-family-management/pkg/v1/family/core"
	"github.com/11SF/go-common/response"
	"github.com/gofiber/fiber/v2"
)

type getFamilyListHandler struct {
	getFamilyList corefamily.GetFamilyList
}

type getFamilyListResponse *[]datamodel.Family

func NewGetFamilyListHandler(getFamilyList corefamily.GetFamilyList) *getFamilyListHandler {
	return &getFamilyListHandler{getFamilyList: getFamilyList}
}

func (h *getFamilyListHandler) Handler(c *fiber.Ctx) error {

	ctx := context.WithValue(c.Context(), "userId", "test")

	res, err := h.getFamilyList(ctx)
	if err != nil {
		slog.Error("failed to get family info", slog.String("error", err.Error()), slog.String("tag", "get family info handler"))
		return response.NewFiberResponseError(c, fiber.StatusOK, err)
	}

	resp := new(getFamilyListResponse)
	*resp = res

	return response.NewFiberResponse(c, response.SuccessCode, resp)
}
