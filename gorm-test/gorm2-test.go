package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ProductA struct {
	gorm.Model
	Code  string
	Price uint
	PCC   uint
}

func init() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&ProductA{})
}
