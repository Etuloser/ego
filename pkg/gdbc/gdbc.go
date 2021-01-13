package gdbc

import (
	"ego/pkg/env"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Product struct
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// GetDb Comment
func GetDb() *gorm.DB {
	envPath := env.GetEnvPath()
	godotenv.Load(envPath)
	databaseURI := os.Getenv("DEV_DATABASE_URI")
	db, err := gorm.Open(mysql.Open(databaseURI), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

// InitDb Comment
func InitDb() {
	db := GetDb()
	// 迁移 schema
	db.AutoMigrate(&Product{})
}

// ProductCreate Comment
func ProductCreate() {
	db := GetDb()
	db.Create(&Product{Code: "D42", Price: 100})
}

// ProductGet Comment
func ProductGet() Product {
	db := GetDb()
	var product Product
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	return product
}
