package main

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 使用GORM

func NewDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("bookstore.db"), &gorm.Config{})
	if err != nil {
		fmt.Printf("sqlite.Open() failed, err: %v", err)
	}

	db.AutoMigrate(&Shelf{}, &Book{})
	return db, nil
}

// defined model

// Shelf 书架
type Shelf struct {
	ID         int64 `gorm:"primaryKey"`
	Theme      string
	Size       int64
	CreateTime time.Time
	UpdateTime time.Time
}

// Book 书
type Book struct {
	ID         int64 `gorm:"primaryKey"`
	Author     string
	Title      string
	ShelfID    int64
	CreateTime time.Time
	UpdateTime time.Time
}
