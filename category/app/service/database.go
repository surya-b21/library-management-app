package service

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// DB global variable
var DB *gorm.DB

// InitDB function to init db
func InitDB() {
	if DB == nil {
		dsn := "host=category_db user=postgres password=password dbname=category_db port=5432 sslmode=disable TimeZone=Asia/Jakarta"
		config := &gorm.Config{
			Logger: logger.New(
				log.New(os.Stderr, "[GORM] ", log.LstdFlags), // io writer
				logger.Config{
					SlowThreshold:             time.Second,   // Slow SQL threshold
					LogLevel:                  logger.Silent, // Log level
					IgnoreRecordNotFoundError: false,         // Ignore ErrRecordNotFound error for logger
					Colorful:                  true,          // Disable color
				},
			),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			DisableForeignKeyConstraintWhenMigrating: true,
		}

		if db, err := gorm.Open(postgres.Open(dsn), config); err == nil {
			DB = db.Debug()
			AutoMigrate(DB)
		}
	}
}
