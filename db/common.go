package db

import (
	"context"
	"fmt"

	"github.com/decus-io/decus-keeper-client/db/dao"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func ConnectDatabase() error {
	var err error
	db, err = gorm.Open(sqlite.Open("keeper.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return fmt.Errorf("open database error: %v", err)
	}
	return db.AutoMigrate(
		&dao.Group{},
		&dao.GroupKeeper{},
		&dao.Withdraw{},
	)
}

func GetDB(ctx context.Context) *gorm.DB {
	return db.WithContext(ctx)
}

func Transaction(ctx context.Context, fc func(context.Context, *gorm.DB) error) error {
	_db := db.WithContext(ctx)

	return _db.Transaction(func(tx *gorm.DB) error {
		return fc(ctx, tx)
	})
}
