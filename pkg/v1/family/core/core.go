package corefamily

import (
	"context"

	"github.com/11SF/core-family-management/pkg/v1/datamodel"
	"github.com/11SF/core-family-management/pkg/v1/family/repository"
)

type service struct {
	db    repository.IFamilyDB
	redis repository.IFamilyRedis
}

type IService interface {
	CreateFamily(ctx context.Context, family *datamodel.Family) (string, error)
	GetFamilyInfo(ctx context.Context, familyId string) (*datamodel.Family, error)
	GetFamilyList(ctx context.Context) (*[]datamodel.Family, error)
	// UpdateFamilyInfo(ctx context.Context, family *datamodel.Family) error
	// UpdateFamilyPrices(ctx context.Context, familyId string, prices []datamodel.Prices) error
	// DeleteFamily(ctx context.Context, familyId string) error
}

func NewService(db repository.IFamilyDB, redis repository.IFamilyRedis) IService {
	return &service{db: db, redis: redis}
}
