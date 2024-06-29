package corefamily

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/11SF/core-family-management/pkg/v1/datamodel"
	"github.com/11SF/go-common/response"
	"github.com/google/uuid"
)

type CreateFamilyFunc func(ctx context.Context, family *datamodel.Family) (string, error)

func (s *service) CreateFamily(ctx context.Context, family *datamodel.Family) (string, error) {

	slog.Info("Starting to create family", slog.String("tag", "core create family"))

	uuid, err := uuid.NewV7()
	if err != nil {
		slog.Error("fail to generate uuid", slog.String("error", err.Error()), slog.String("tag", "core create family"))
		return "", response.NewError("CFM550", err.Error())
	}
	family.ID = uuid.String()

	priceByte, err := json.Marshal(family.Prices)
	if err != nil {
		slog.Error("fail to marshal prices", slog.String("error", err.Error()), slog.String("tag", "core create family"))
		return "", response.NewError("CFM551", err.Error())
	}
	family.PricesString = string(priceByte)
	family.CreatedAt = time.Now()
	family.CreatedBy = ctx.Value("userId").(string)

	err = s.db.CreateFamily(ctx, family)
	if err != nil {
		slog.Error("fail to create family in db", slog.String("error", err.Error()), slog.String("tag", "core create family"))
		return "", response.NewError("CFM552", err.Error())
	}

	slog.Info("Create family success", slog.String("tag", "core create family"))

	return family.ID, nil
}
