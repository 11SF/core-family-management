package repository

import (
	"context"

	"github.com/11SF/core-family-management/pkg/v1/datamodel"
)

type IFamilyDB interface {
	CreateFamily(ctx context.Context, family *datamodel.Family) error
	GetFamilyById(ctx context.Context, familyId string) (*datamodel.Family, error)
	GetFamilyList(ctx context.Context, userId string) (*[]datamodel.Family, error)
	UpdateFamilyInfo(ctx context.Context, family *datamodel.Family, userId string) error
	DeleteFamily(ctx context.Context, familyId string, userId string) error
}

type IFamilyRedis interface {
	SaveFamily(ctx context.Context, family *datamodel.Family) error
	GetFamilyById(ctx context.Context, familyId string) (*datamodel.Family, error)
	DeleteFamily(ctx context.Context, familyId string) error
	SaveFamilyList(ctx context.Context, families []datamodel.Family, userId string) error
	GetFamilyList(ctx context.Context, userId string) (*[]datamodel.Family, error)
	DeleteFamilyList(ctx context.Context, userId string) error
}
