package coregoods

import "github.com/11SF/core-family-management/pkg/v1/goods/repository"

type service struct {
	db repository.IGoodsDB
}

func NewService(db repository.IGoodsDB) *service {
	return &service{db: db}
}
