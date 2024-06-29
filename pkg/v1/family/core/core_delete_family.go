package corefamily

import (
	"context"
	"log/slog"

	"github.com/11SF/go-common/response"
)

type DeleteFamily func(ctx context.Context, familyId string) error

func (s *service) DeleteFamily(ctx context.Context, familyId string) error {

	slog.Info("Starting to delete family", slog.String("family", familyId), slog.String("tag", "core delete family"))

	userId := ctx.Value("userId").(string)

	err := s.db.DeleteFamily(ctx, familyId, userId)
	if err != nil {
		slog.Error("fail to delete family in db", slog.String("error", err.Error()), slog.String("tag", "core delete family"))
		return response.NewError("CFM552", err.Error())
	}

	err = s.redis.DeleteFamilyList(ctx, userId)
	if err != nil {
		slog.Error("fail to delete family list in redis", slog.String("error", err.Error()), slog.String("tag", "core update family info"))
		return response.NewError("CFM553", err.Error())
	}

	err = s.redis.DeleteFamily(ctx, familyId)
	if err != nil {
		slog.Error("fail to delete family in redis", slog.String("error", err.Error()), slog.String("tag", "core update family info"))
		return response.NewError("CFM554", err.Error())
	}

	slog.Info("Success to delete family", slog.String("family", familyId), slog.String("tag", "core delete family"))

	return nil
}
