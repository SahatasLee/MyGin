package db

import (
	"context"
	"fmt"
	"myGin/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()

	fmt.Print(sql + "\n ================================= \n")
}

func DatabaseInit() *gorm.DB {
	dsn := "host=10.111.0.124 user=admin password=psltest dbname=postgresdb port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: &SqlLogger{},
		DryRun: false,
	})

	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(
		models.User{},
		models.Book{},
	)

	return db
}
