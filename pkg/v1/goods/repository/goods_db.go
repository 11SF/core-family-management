package repository

import (
	"context"

	"github.com/11SF/core-family-management/pkg/v1/datamodel"
	"github.com/uptrace/bun"
)

type goodsDB struct {
	db *bun.DB
}

type IGoodsDB interface {
	GetGoodsByID(ctx context.Context, id string) (*datamodel.Goods, error)
}

func NewGoodsDB(db *bun.DB) *goodsDB {
	return &goodsDB{db: db}
}

func (r *goodsDB) GetGoodsByID(ctx context.Context, id string) (*datamodel.Goods, error) {
	goods := new(datamodel.Goods)
	err := r.db.NewSelect().Model(goods).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return goods, nil
}
