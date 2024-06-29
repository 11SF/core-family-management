package corefamily

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/11SF/core-family-management/pkg/v1/datamodel"
	"github.com/11SF/go-common/response"
)

type GetFamilyInfo func(ctx context.Context, familyId string) (*datamodel.Family, error)

func (s *service) GetFamilyInfo(ctx context.Context, familyId string) (*datamodel.Family, error) {

	slog.Info("Starting to get family info", slog.String("familyId", familyId), slog.String("tag", "core get family info"))

	userId := ctx.Value("userId").(string)
	if userId == "" {
		return nil, response.NewError("CFM444", "userId is required")
	}

	familyInfo, err := s.redis.GetFamilyById(ctx, familyId)
	if err != nil {
		slog.Error("failed to get family info from redis try to get from db", slog.String("error", err.Error()), slog.String("tag", "core get family info"))

		familyInfo, err = s.db.GetFamilyById(ctx, familyId)
		if err != nil {
			return nil, response.NewError("CFM560", err.Error())
		}

		pricePayload := new([]datamodel.Prices)
		err = json.Unmarshal([]byte(familyInfo.PricesString), pricePayload)
		if err != nil {
			slog.Error("failed to unmarshal prices", slog.String("error", err.Error()), slog.String("tag", "core get family info"))
			return nil, response.NewError("CFM561", err.Error())
		}
		familyInfo.Prices = pricePayload

		err = s.redis.SaveFamily(ctx, familyInfo)
		if err != nil {
			slog.Error("failed to save family info to redis", slog.String("error", err.Error()), slog.String("tag", "core get family info"))
		}
	}

	if familyInfo.CreatedBy != userId {
		return nil, response.NewError("CFM403", "Permission denied")
	}

	slog.Info("Success to get family info", slog.String("familyId", familyId), slog.String("tag", "core get family info"))

	return familyInfo, nil
}
