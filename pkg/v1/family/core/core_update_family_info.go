package corefamily

import (
	"context"
	"log/slog"
	"time"

	"github.com/11SF/core-family-management/pkg/v1/datamodel"
	"github.com/11SF/go-common/response"
)

type UpdateFamilyInfo func(ctx context.Context, family *datamodel.Family) error

func (s *service) UpdateFamilyInfo(ctx context.Context, family *datamodel.Family) error {

	slog.Info("Starting to update family info", slog.String("family", family.ID), slog.String("tag", "core update family info"))

	userId := ctx.Value("userId").(string)

	family.UpdatedAt = time.Now()
	family.UpdatedBy = userId

	err := s.db.UpdateFamilyInfo(ctx, family, userId)
	if err != nil {
		slog.Error("fail to update family info in db", slog.String("error", err.Error()), slog.String("tag", "core update family info"))
		return response.NewError("CFM552", err.Error())
	}

	err = s.redis.DeleteFamilyList(ctx, userId)
	if err != nil {
		slog.Error("fail to delete family list in redis", slog.String("error", err.Error()), slog.String("tag", "core update family info"))
		return response.NewError("CFM553", err.Error())
	}

	err = s.redis.DeleteFamily(ctx, family.ID)
	if err != nil {
		slog.Error("fail to delete family in redis", slog.String("error", err.Error()), slog.String("tag", "core update family info"))
		return response.NewError("CFM554", err.Error())
	}

	slog.Info("Success to update family info", slog.String("family", family.ID), slog.String("tag", "core update family info"))
	return nil
}
