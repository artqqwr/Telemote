package repository

import "gorm.io/gorm"

type Repository interface {
}

type repository struct {
}

func New(db *gorm.DB) Repository {
	return &repository{}
}
