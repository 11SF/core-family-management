package coregoods

import (
	"context"

	"github.com/11SF/core-family-management/pkg/v1/datamodel"
)

type GetGoodsByIdFunc func(ctx context.Context, id string) (*datamodel.Goods, error)

func (s *service) GetGoodsByID(ctx context.Context, id string) (*datamodel.Goods, error) {
	return s.db.GetGoodsByID(ctx, id)
}
