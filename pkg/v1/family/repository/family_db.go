package repository

import (
	"context"

	"github.com/11SF/core-family-management/pkg/v1/datamodel"
	"gorm.io/gorm"
)

type familyDB struct {
	db *gorm.DB
}

func NewFamilyDB(db *gorm.DB) IFamilyDB {
	err := db.AutoMigrate(&datamodel.Family{})
	if err != nil {
		panic(err)
	}

	return &familyDB{db: db}
}

func (r *familyDB) CreateFamily(ctx context.Context, family *datamodel.Family) error {

	err := r.db.Create(family).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *familyDB) GetFamilyById(ctx context.Context, familyId string) (*datamodel.Family, error) {
	var family datamodel.Family
	err := r.db.Where("id = ?", familyId).First(&family).Error
	if err != nil {
		return nil, err
	}

	return &family, nil
}

func (r *familyDB) GetFamilyList(ctx context.Context, userId string) ([]datamodel.Family, error) {
	var families []datamodel.Family
	err := r.db.Where("user_id = ?", userId).Find(&families).Error
	if err != nil {
		return nil, err
	}

	return families, nil
}

func (r *familyDB) UpdateFamilyInfo(ctx context.Context, family *datamodel.Family) error {
	err := r.db.Model(family).Updates(family).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *familyDB) UpdateFamilyPrices(ctx context.Context, familyId string, prices []datamodel.Prices) error {
	err := r.db.Model(&datamodel.Prices{}).Where("family_id = ?", familyId).Updates(prices).Error
	if err != nil {
		return err
	}
	return nil
}
