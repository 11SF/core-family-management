package repository

import (
	"context"
	"errors"
	"time"

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

func (r *familyDB) GetFamilyList(ctx context.Context, userId string) (*[]datamodel.Family, error) {
	families := new([]datamodel.Family)
	err := r.db.Where("created_by = ?", userId).Find(&families).Error
	if err != nil {
		return nil, err
	}

	return families, nil
}

func (r *familyDB) UpdateFamilyInfo(ctx context.Context, family *datamodel.Family, userId string) error {
	result := r.db.Model(family).Where("created_by = ?", userId).Updates(family)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("family not found")
	}

	return nil
}

func (r *familyDB) DeleteFamily(ctx context.Context, familyId string, userId string) error {
	result := r.db.Model(&datamodel.Family{}).
		Where("id = ? AND created_by = ?", familyId, userId).
		Updates(datamodel.Family{
			DeletedAt: gorm.DeletedAt{Time: time.Now(), Valid: true},
			DeletedBy: userId,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("family not found")
	}

	return nil
}
