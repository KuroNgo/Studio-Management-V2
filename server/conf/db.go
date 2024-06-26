package conf

import (
	"co-studio-e-commerce/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DbDefault *gorm.DB

// GetDB sử dụng singleton pattern để tạo một connection duy nhất đến database
// khi ứng dụng lớn hơn thì không nên sử dụng singleton pattern
// thay vào đó nên sử dụng connection pool
func (a *App) GetDB() *gorm.DB {
	if DbDefault == nil {
		return a.initDB()
	}
	return DbDefault
}

func (a *App) initDB() *gorm.DB {
	//  Tạo chuỗi kết nối đến PostgreSQL
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Cfg.DBHost, Cfg.Port, Cfg.DBUser, Cfg.DBPassword, Cfg.DBName)

	// Mở kết nối đến PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	db.Exec("CREATE SCHEMA IF NOT EXISTS public")
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	err = db.AutoMigrate(
		&model.User{},
		&model.Product{},
		&model.Image{},
		&model.Category{},
		&model.Order{},
		&model.OrderDetail{},
		&model.Admin{},
		&model.ActivityLog{},
		&model.LoginSession{},
		&model.PaymentInformation{},
		&model.Post{},
		&model.ShoppingCart{},
		&model.Reviews{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	return db
}
