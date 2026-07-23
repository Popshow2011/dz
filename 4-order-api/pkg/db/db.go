package db

import (
	"dz/4-order-api/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.Config
}

func NewDb(conf *configs.Config) *Db {
	db, err := gorm.Open(postgres.Open(conf.Db.DSN), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return &Db{db.Config}
}
