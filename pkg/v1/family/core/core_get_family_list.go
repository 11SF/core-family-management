package corefamily

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/11SF/core-family-management/pkg/v1/datamodel"
	"github.com/11SF/go-common/response"
)

type GetFamilyList func(ctx context.Context) (*[]datamodel.Family, error)

func (s *service) GetFamilyList(ctx context.Context) (*[]datamodel.Family, error) {

	userId := ctx.Value("userId").(string)
	if userId == "" {
		return nil, response.NewError("CFM403", "Permission Denied")
	}

	slog.Info("Starting to get family list", slog.String("userId", userId), slog.String("tag", "core get family list"))

	families, err := s.redis.GetFamilyList(ctx, userId)
	if err != nil {
		slog.Error("failed to get family list from redis try to get from db", slog.String("err", err.Error()), slog.String("tag", "core get family list"))

		families, err = s.db.GetFamilyList(ctx, userId)
		if err != nil {
			slog.Error("failed to get family list from db", slog.String("err", err.Error()), slog.String("tag", "core get family list"))
			return nil, response.NewError("CFM560", "Internal Server Error")
		}

		for index, family := range *families {
			pricePayload := new([]datamodel.Prices)
			err = json.Unmarshal([]byte(family.PricesString), pricePayload)
			if err != nil {
				slog.Error("failed to unmarshal prices", slog.String("error", err.Error()), slog.String("tag", "core get family info"))
				return nil, response.NewError("CFM561", err.Error())
			}
			(*families)[index].Prices = pricePayload
		}

		err = s.redis.SaveFamilyList(ctx, *families, userId)
		if err != nil {
			slog.Error("failed to save family list to redis", slog.String("err", err.Error()), slog.String("tag", "core get family list"))
		}
	}

	slog.Info("Success to get family list", slog.String("userId", userId), slog.String("tag", "core get family list"))

	return families, nil
}
