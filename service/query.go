package service

import "gorm.io/gorm"

type Query interface {
}

type query struct {
	db *gorm.DB
}

func NewQuery(db *gorm.DB) *query {
	return &query{db}
}
